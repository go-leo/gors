package gors

import (
	"context"
	"sync/atomic"
)

type codeKey struct{}

func InjectStatusCode(ctx context.Context, code int) context.Context {
	codePtr := new(int32)
	*codePtr = int32(code)
	return context.WithValue(ctx, codeKey{}, codePtr)
}

func SetStatusCode(ctx context.Context, code int) {
	v := ctx.Value(codeKey{})
	codePtr, _ := v.(*int32)
	atomic.SwapInt32(codePtr, int32(code))
}

func GetStatusCode(ctx context.Context) int {
	v := ctx.Value(codeKey{})
	codePtr, _ := v.(*int32)
	return int(atomic.LoadInt32(codePtr))
}
