package status

import (
	"net/http"

	gcodes "google.golang.org/grpc/codes"
)

const (
	// ClientClosed is non-standard http status code,
	// which defined by nginx.
	// https://httpstatus.in/499/
	ClientClosed = 499
)

// Converter is a status converter.
type Converter interface {
	// ToGRPCCode converts an HTTP error code into the corresponding gRPC response status.
	ToGRPCCode(code int) gcodes.Code

	// FromGRPCCode converts a gRPC error code into the corresponding HTTP response status.
	FromGRPCCode(code gcodes.Code) int
}

type statusConverter struct{}

// DefaultConverter default converter.
var DefaultConverter Converter = statusConverter{}

// ToGRPCCode converts a HTTP error code into the corresponding gRPC response status.
// See: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
func (c statusConverter) ToGRPCCode(code int) gcodes.Code {
	switch code {
	case http.StatusOK:
		return gcodes.OK
	case http.StatusBadRequest:
		return gcodes.InvalidArgument
	case http.StatusUnauthorized:
		return gcodes.Unauthenticated
	case http.StatusForbidden:
		return gcodes.PermissionDenied
	case http.StatusNotFound:
		return gcodes.NotFound
	case http.StatusConflict:
		return gcodes.Aborted
	case http.StatusTooManyRequests:
		return gcodes.ResourceExhausted
	case http.StatusInternalServerError:
		return gcodes.Internal
	case http.StatusNotImplemented:
		return gcodes.Unimplemented
	case http.StatusServiceUnavailable:
		return gcodes.Unavailable
	case http.StatusGatewayTimeout:
		return gcodes.DeadlineExceeded
	case ClientClosed:
		return gcodes.Canceled
	}

	return gcodes.Unknown
}

// FromGRPCCode converts a gRPC error code into the corresponding HTTP response status.
// See: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
func (c statusConverter) FromGRPCCode(code gcodes.Code) int {
	switch code {
	case gcodes.OK:
		return http.StatusOK
	case gcodes.Canceled:
		return ClientClosed
	case gcodes.Unknown:
		return http.StatusInternalServerError
	case gcodes.InvalidArgument:
		return http.StatusBadRequest
	case gcodes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case gcodes.NotFound:
		return http.StatusNotFound
	case gcodes.AlreadyExists:
		return http.StatusConflict
	case gcodes.PermissionDenied:
		return http.StatusForbidden
	case gcodes.Unauthenticated:
		return http.StatusUnauthorized
	case gcodes.ResourceExhausted:
		return http.StatusTooManyRequests
	case gcodes.FailedPrecondition:
		return http.StatusBadRequest
	case gcodes.Aborted:
		return http.StatusConflict
	case gcodes.OutOfRange:
		return http.StatusBadRequest
	case gcodes.Unimplemented:
		return http.StatusNotImplemented
	case gcodes.Internal:
		return http.StatusInternalServerError
	case gcodes.Unavailable:
		return http.StatusServiceUnavailable
	case gcodes.DataLoss:
		return http.StatusInternalServerError
	}

	return http.StatusInternalServerError
}

// ToGRPCCode converts an HTTP error code into the corresponding gRPC response status.
func ToGRPCCode(code int) gcodes.Code {
	return DefaultConverter.ToGRPCCode(code)
}

// FromGRPCCode converts a gRPC error code into the corresponding HTTP response status.
func FromGRPCCode(code gcodes.Code) int {
	return DefaultConverter.FromGRPCCode(code)
}
