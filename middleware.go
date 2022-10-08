package gors

import (
	"context"
	"net/http"
)

type ServerInfo struct {
	Server         any
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}

// Handler defines the handler invoked by ServerInterceptor to complete the normal
// execution of a RPC.
//
// If a Handler returns an error, it should either be produced by the status package,
// or be one of the context errors. Otherwise, gRPC will use codes.Unknown as the status
// code and err.Error() as the status message of the RPC.
type Handler func(ctx context.Context, req any) (resp any, err error)

// ServerInterceptor provides a hook to intercept the execution of a RPC on the server. info
// contains all the information of this RPC the interceptor can operate on. And handler is the wrapper
// of the service method implementation. It is the responsibility of the interceptor to invoke handler
// to complete the RPC.
type ServerInterceptor func(ctx context.Context, req any, info *ServerInfo, handler Handler) (resp any, err error)

// ChainServerInterceptors chains all server interceptors into one.
func ChainServerInterceptors(interceptors ...ServerInterceptor) ServerInterceptor {
	var chainedInt ServerInterceptor
	if len(interceptors) == 0 {
		chainedInt = nil
	} else if len(interceptors) == 1 {
		chainedInt = interceptors[0]
	} else {
		chainedInt = chainUnaryInterceptors(interceptors)
	}
	return chainedInt
}

func chainUnaryInterceptors(interceptors []ServerInterceptor) ServerInterceptor {
	return func(ctx context.Context, req any, info *ServerInfo, handler Handler) (any, error) {
		// the struct ensures the variables are allocated together, rather than separately, since we
		// know they should be garbage collected together. This saves 1 allocation and decreases
		// time/call by about 10% on the microbenchmark.
		var state struct {
			i    int
			next Handler
		}
		state.next = func(ctx context.Context, req any) (any, error) {
			if state.i == len(interceptors)-1 {
				return interceptors[state.i](ctx, req, info, handler)
			}
			state.i++
			return interceptors[state.i-1](ctx, req, info, state.next)
		}
		return state.next(ctx, req)
	}
}
