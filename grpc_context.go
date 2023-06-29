package gors

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/go-leo/gox/stringx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/textproto"
	"strconv"
	"strings"
	"sync"
	"time"
)

// DefaultContextTimeout is used for gRPC call context.WithTimeout whenever a Grpc-Timeout inbound
// header isn't present. If the value is 0 the sent `context` will not have a timeout.
var DefaultContextTimeout = 0 * time.Second

// MetadataTrailerPrefix is prepended to gRPC metadata as it is converted to
// HTTP headers in a response handled by grpc-gateway
const MetadataTrailerPrefix = "Grpc-Trailer-"

const metadataGrpcTimeout = "Grpc-Timeout"
const metadataHeaderBinarySuffix = "-Bin"

const xForwardedFor = "X-Forwarded-For"
const xForwardedHost = "X-Forwarded-Host"
const authorization = "Authorization"

// GorsMetadataPrefix is prepended to permanent HTTP header keys when added to the gRPC context.
const GorsMetadataPrefix = "gors-"

// MetadataHeaderPrefix is the http prefix that represents custom metadata
// parameters to or from a gRPC call.
const MetadataHeaderPrefix = "Grpc-Metadata-"

var _ grpc.ServerTransportStream = new(ServerTransportStream)

// ServerTransportStream implements grpc.ServerTransportStream.
// It should only be used by the generated files to support grpc.SendHeader
// outside of gRPC server use.
type ServerTransportStream struct {
	mu      sync.RWMutex
	header  metadata.MD
	trailer metadata.MD
	method  string
}

func NewServerTransportStream(method string) *ServerTransportStream {
	return &ServerTransportStream{method: method}
}

func (s *ServerTransportStream) Method() string {
	return s.method
}

// Header returns the header metadata of the stream.
func (s *ServerTransportStream) Header() metadata.MD {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.header.Copy()
}

// SetHeader sets the header metadata.
func (s *ServerTransportStream) SetHeader(md metadata.MD) error {
	if md.Len() == 0 {
		return nil
	}

	s.mu.Lock()
	s.header = metadata.Join(s.header, md)
	s.mu.Unlock()
	return nil
}

func (s *ServerTransportStream) SendHeader(md metadata.MD) error {
	return s.SetHeader(md)
}

// Trailer returns the cached trailer metadata.
func (s *ServerTransportStream) Trailer() metadata.MD {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.trailer.Copy()
}

// SetTrailer sets the trailer metadata.
func (s *ServerTransportStream) SetTrailer(md metadata.MD) error {
	if md.Len() == 0 {
		return nil
	}

	s.mu.Lock()
	s.trailer = metadata.Join(s.trailer, md)
	s.mu.Unlock()
	return nil
}

// NewGRPCContext adds context information such as metadata from the request.
// At a minimum, the RemoteAddr is included in the fashion of "X-Forwarded-For",
// except that the forwarded destination is not another HTTP service but rather
// a gRPC service.
func NewGRPCContext(
	ctx context.Context,
	headerMatcher func(key string) (string, bool),
	metadataAnnotators []func(ctx context.Context) metadata.MD,
) (context.Context, error) {
	c := FromContext(ctx)

	timeout := DefaultContextTimeout
	if tm := c.GetHeader(metadataGrpcTimeout); tm != "" {
		var err error
		timeout, err = timeoutDecode(tm)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid grpc-timeout: %s", tm)
		}
	}

	var pairs []string

	if host := c.GetHeader(xForwardedHost); host != "" {
		pairs = append(pairs, strings.ToLower(xForwardedHost), host)
	} else if c.Request.Host != "" {
		pairs = append(pairs, strings.ToLower(xForwardedHost), c.Request.Host)
	}

	if remoteIP := c.RemoteIP(); remoteIP != "" {
		if fwd := c.GetHeader(xForwardedFor); stringx.IsNotBlank(fwd) {
			pairs = append(pairs, strings.ToLower(xForwardedFor), fmt.Sprintf("%s, %s", fwd, remoteIP))
		} else {
			pairs = append(pairs, strings.ToLower(xForwardedFor), remoteIP)
		}
	}

	if auth := c.GetHeader(authorization); stringx.IsNotBlank(auth) {
		pairs = append(pairs, "authorization", auth)
	}

	if headerMatcher == nil {
		headerMatcher = defaultIncomingHeaderMatcher
	}
	for key, values := range c.Request.Header {
		key = textproto.CanonicalMIMEHeaderKey(key)
		switch key {
		case metadataGrpcTimeout:
			continue
		case xForwardedHost:
			continue
		case xForwardedFor:
			continue
		case authorization:
			continue
		default:
			h, ok := headerMatcher(key)
			if !ok {
				continue
			}
			if !isValidGRPCMetadataKey(h) {
				continue
			}
			for _, value := range values {
				if strings.HasSuffix(key, metadataHeaderBinarySuffix) {
					b, err := decodeBinHeader(value)
					if err != nil {
						return nil, status.Errorf(codes.InvalidArgument, "invalid binary header %s: %s", key, err)
					}
					pairs = append(pairs, h, string(b))
					continue
				}
				if !isValidGRPCMetadataTextValue(value) {
					continue
				}
			}
		}
	}

	if timeout > 0 {
		ctx, _ = context.WithTimeout(ctx, timeout) //nolint:govet  the context outlives this function
	}

	var md metadata.MD
	if len(pairs) > 0 {
		md = metadata.Join(md, metadata.Pairs(pairs...))
	}
	for _, mda := range metadataAnnotators {
		md = metadata.Join(md, mda(ctx))
	}
	if md.Len() > 0 {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	return ctx, nil
}

func timeoutDecode(s string) (time.Duration, error) {
	size := len(s)
	if size < 2 {
		return 0, fmt.Errorf("timeout string is too short: %q", s)
	}
	d, ok := timeoutUnitToDuration(s[size-1])
	if !ok {
		return 0, fmt.Errorf("timeout unit is not recognized: %q", s)
	}
	t, err := strconv.ParseInt(s[:size-1], 10, 64)
	if err != nil {
		return 0, err
	}
	return d * time.Duration(t), nil
}

func timeoutUnitToDuration(u uint8) (d time.Duration, ok bool) {
	switch u {
	case 'H':
		return time.Hour, true
	case 'M':
		return time.Minute, true
	case 'S':
		return time.Second, true
	case 'm':
		return time.Millisecond, true
	case 'u':
		return time.Microsecond, true
	case 'n':
		return time.Nanosecond, true
	default:
		return
	}
}

// defaultIncomingHeaderMatcher is used to pass http request headers to/from gRPC context. This adds permanent HTTP header
// keys (as specified by the IANA, e.g: Accept, Cookie, Host) to the gRPC metadata with the grpcgateway- prefix. If you want to know which headers are considered permanent, you can view the isPermanentHTTPHeader function.
// HTTP headers that start with 'Grpc-Metadata-' are mapped to gRPC metadata after removing the prefix 'Grpc-Metadata-'.
// Other headers are not added to the gRPC metadata.
func defaultIncomingHeaderMatcher(key string) (string, bool) {
	switch key = textproto.CanonicalMIMEHeaderKey(key); {
	case isPermanentHTTPHeader(key):
		return GorsMetadataPrefix + key, true
	case strings.HasPrefix(key, MetadataHeaderPrefix):
		return key[len(MetadataHeaderPrefix):], true
	}
	return "", false
}

// isPermanentHTTPHeader checks whether hdr belongs to the list of
// permanent request headers maintained by IANA.
// http://www.iana.org/assignments/message-headers/message-headers.xml
func isPermanentHTTPHeader(hdr string) bool {
	switch hdr {
	case
		"Accept",
		"Accept-Charset",
		"Accept-Language",
		"Accept-Ranges",
		"Authorization",
		"Cache-Control",
		"Content-Type",
		"Cookie",
		"Date",
		"Expect",
		"From",
		"Host",
		"If-Match",
		"If-Modified-Since",
		"If-None-Match",
		"If-Schedule-Tag-Match",
		"If-Unmodified-Since",
		"Max-Forwards",
		"Origin",
		"Pragma",
		"Referer",
		"User-Agent",
		"Via",
		"Warning":
		return true
	}
	return false
}

func isValidGRPCMetadataKey(key string) bool {
	// Must be a valid gRPC "Header-Name" as defined here:
	//   https://github.com/grpc/grpc/blob/4b05dc88b724214d0c725c8e7442cbc7a61b1374/doc/PROTOCOL-HTTP2.md
	// This means 0-9 a-z _ - .
	// Only lowercase letters are valid in the wire protocol, but the client library will normalize
	// uppercase ASCII to lowercase, so uppercase ASCII is also acceptable.
	bytes := []byte(key) // gRPC validates strings on the byte level, not Unicode.
	for _, ch := range bytes {
		validLowercaseLetter := ch >= 'a' && ch <= 'z'
		validUppercaseLetter := ch >= 'A' && ch <= 'Z'
		validDigit := ch >= '0' && ch <= '9'
		validOther := ch == '.' || ch == '-' || ch == '_'
		if !validLowercaseLetter && !validUppercaseLetter && !validDigit && !validOther {
			return false
		}
	}
	return true
}

func decodeBinHeader(v string) ([]byte, error) {
	if len(v)%4 == 0 {
		// Input was padded, or padding was not necessary.
		return base64.StdEncoding.DecodeString(v)
	}
	return base64.RawStdEncoding.DecodeString(v)
}

func isValidGRPCMetadataTextValue(textValue string) bool {
	// Must be a valid gRPC "ASCII-Value" as defined here:
	//   https://github.com/grpc/grpc/blob/4b05dc88b724214d0c725c8e7442cbc7a61b1374/doc/PROTOCOL-HTTP2.md
	// This means printable ASCII (including/plus spaces); 0x20 to 0x7E inclusive.
	bytes := []byte(textValue) // gRPC validates strings on the byte level, not Unicode.
	for _, ch := range bytes {
		if ch < 0x20 || ch > 0x7E {
			return false
		}
	}
	return true
}
