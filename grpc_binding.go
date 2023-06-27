package gors

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"net/http"
	"net/textproto"
	"strings"
)

var OutgoingHeaderMatcher = DefaultOutgoingHeaderMatcher

var HandleGRPCError = DefaultHandleGRPCError

func ShouldBindWith(c *gin.Context, req any, tag string, fns ...func(c *gin.Context, req any, tag string) error) error {
	for _, fn := range fns {
		if err := fn(c, req, tag); err != nil {
			return err
		}
	}
	return nil
}

func UriBindingWith(c *gin.Context, req any, tag string) error {
	m := make(map[string][]string)
	for _, v := range c.Params {
		m[v.Key] = []string{v.Value}
	}
	return binding.MapFormWithTag(req, m, tag)
}

func QueryBindingWith(c *gin.Context, req any, tag string) error {
	return binding.MapFormWithTag(req, c.Request.URL.Query(), tag)
}

func HeaderBindingWith(c *gin.Context, req any, tag string) error {
	return binding.MapFormWithTag(req, c.Request.Header, tag)
}

func FormBindingWith(c *gin.Context, req any, tag string) error {
	if err := c.Request.ParseForm(); err != nil {
		return err
	}
	const defaultMemory = 32 << 20
	if err := c.Request.ParseMultipartForm(defaultMemory); err != nil && !errors.Is(err, http.ErrNotMultipart) {
		return err
	}
	return binding.MapFormWithTag(req, c.Request.Form, tag)
}

func FormPostBindingWith(c *gin.Context, req any, tag string) error {
	if err := c.Request.ParseForm(); err != nil {
		return err
	}
	return binding.MapFormWithTag(req, c.Request.PostForm, tag)
}

func JSONBindingWith(c *gin.Context, req any, tag string) error {
	return c.ShouldBindWith(req, binding.JSON)
}

func ProtoBufBindingWith(c *gin.Context, req any, tag string) error {
	return c.ShouldBindWith(req, binding.ProtoBuf)
}

func MsgPackBindingWith(c *gin.Context, req any, tag string) error {
	return c.ShouldBindWith(req, binding.MsgPack)
}

func CustomBindingWith(c *gin.Context, req any, tag string) error {
	customBinding, ok := req.(Binding)
	if !ok {
		return nil
	}
	return customBinding.Bind(c)
}

func DefaultHandleGRPCError(c *gin.Context, err error) {
	var customStatus *HttpError
	if errors.As(err, &customStatus) {
		err = customStatus.err
	}

	s := status.Convert(err)
	pb := s.Proto()

	c.Writer.Header().Del("Trailer")
	c.Writer.Header().Del("Transfer-Encoding")

	if s.Code() == codes.Unauthenticated {
		c.Writer.Header().Set("WWW-Authenticate", s.Message())
	}

	md, ok := GRPCMetadataFromContext(c.Request.Context())
	if !ok {
		grpclog.Infof("Failed to extract GRPCMetadata from context")
	}

	handleForwardResponseGRPCMetadata(c, md)

	// RFC 7230 https://tools.ietf.org/html/rfc7230#section-4.1.2
	// Unless the request includes a TE header field indicating "trailers"
	// is acceptable, as described in Section 4.3, a server SHOULD NOT
	// generate trailer fields that it believes are necessary for the user
	// agent to receive.
	doForwardTrailers := requestAcceptsTrailers(c)

	if doForwardTrailers {
		handleForwardResponseTrailerHeader(c, md)
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
	}

	st := HTTPStatusFromCode(s.Code())
	if customStatus != nil {
		st = customStatus.statusCode
	}

	c.String(st, pb.Message)

	if doForwardTrailers {
		handleForwardResponseTrailer(c, md)
	}
}

func handleForwardResponseGRPCMetadata(c *gin.Context, md GRPCMetadata) {
	for k, vs := range md.HeaderMD {
		if h, ok := OutgoingHeaderMatcher(k); ok {
			for _, v := range vs {
				c.Writer.Header().Add(h, v)
			}
		}
	}
}

func requestAcceptsTrailers(c *gin.Context) bool {
	te := c.GetHeader("TE")
	return strings.Contains(strings.ToLower(te), "trailers")
}

func handleForwardResponseTrailerHeader(c *gin.Context, md GRPCMetadata) {
	for k := range md.TrailerMD {
		tKey := textproto.CanonicalMIMEHeaderKey(fmt.Sprintf("%s%s", MetadataTrailerPrefix, k))
		c.Writer.Header().Add("Trailer", tKey)
	}
}

func handleForwardResponseTrailer(c *gin.Context, md GRPCMetadata) {
	for k, vs := range md.TrailerMD {
		tKey := fmt.Sprintf("%s%s", MetadataTrailerPrefix, k)
		for _, v := range vs {
			c.Writer.Header().Add(tKey, v)
		}
	}
}

func DefaultOutgoingHeaderMatcher(key string) (string, bool) {
	return fmt.Sprintf("%s%s", MetadataHeaderPrefix, key), true
}

// HTTPStatusFromCode converts a gRPC error code into the corresponding HTTP response status.
// See: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
func HTTPStatusFromCode(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return 499
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		// Note, this deliberately doesn't translate to the similarly named '412 Precondition Failed' HTTP response status.
		return http.StatusBadRequest
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	default:
		grpclog.Infof("Unknown gRPC error code: %v", code)
		return http.StatusInternalServerError
	}
}
