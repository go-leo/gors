package gors

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试Error解析
func TestErrorIs(t *testing.T) {
	// 同一对象
	testerr := Error{
		StatusCode: http.StatusBadRequest,
		Code:       10002,
		Message:    "test",
	}
	var err error = testerr
	assert.True(t, testerr.Is(err))

	// 不同对象，相同code
	var err2 error = Error{
		StatusCode: http.StatusBadRequest,
		Code:       10002,
		Message:    "test2",
	}
	assert.True(t, testerr.Is(err2))

	// warpper后的对象，相同code
	warperr := fmt.Errorf("wrap it: %w", err2)
	assert.True(t, testerr.Is(warperr))

	// 指针和值对象，相同code
	gerr := &Error{
		StatusCode: http.StatusBadRequest,
		Code:       10002,
		Message:    "test2",
	}
	assert.True(t, testerr.Is(gerr))
	assert.True(t, gerr.Is(testerr))

	// 不同code的对象
	var err4 error = Error{
		StatusCode: http.StatusBadRequest,
		Code:       10003,
		Message:    "test3",
	}
	assert.False(t, testerr.Is(err4))
}

// 测试Error解析
func TestFromError(t *testing.T) {
	var err error = Error{
		StatusCode: http.StatusBadRequest,
		Code:       10002,
		Message:    "test",
	}

	// 解析*Error到*Error
	e := FromError(err)
	assert.Equal(t, err, e)

	// 解析*status.Status到*Error
	grpcerr := e.GRPCStatus()
	e2 := FromError(grpcerr.Err())
	assert.Equal(t, err, e2)
}

// 测试兼容性 标准库 errors.Is、As、Unwarp
func TestCompatibility(t *testing.T) {
	err := Error{
		StatusCode: http.StatusBadRequest,
		Code:       10002,
		Message:    "test",
	}

	type args struct {
		err    error
		target error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "std errors is、as、unwrap",
			args: args{
				err:    fmt.Errorf("wrap it: %w", err),
				target: err,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, errors.Is(tt.args.err, tt.args.target), tt.want)
			v := Error{}
			assert.Equal(t, errors.As(tt.args.err, &v), tt.want)
			assert.Equal(t, errors.Unwrap(tt.args.err), tt.args.target)
		})
	}
}
