package gors

import (
	"errors"
	"fmt"
	"github.com/go-leo/gox/errorx"
	"io"
	"regexp"

	"github.com/go-leo/gors/internal/pkg/status"
	"github.com/go-leo/gox/convx"
	gstatus "google.golang.org/grpc/status"
)

// 全局默认错误码，可以自定义覆盖
var (
	UnknownStatusCode = 500
	UnknownCode       = 100001
	UnknownMessage    = "An internal server error occurred"
)

var msgRegExp = regexp.MustCompile("^gors.Error, StatusCode: (\\d+), Code: (\\d+), Message: (.+)$")

// Status 代表业务状态信息.
type Status struct {
	Code    int    `json:"code,omitempty" yaml:"code,omitempty" xml:"code,omitempty" toml:"code,omitempty" codec:"code,omitempty" mapstructure:"code,omitempty"`
	Message string `json:"message,omitempty" yaml:"message,omitempty" xml:"message,omitempty" toml:"message,omitempty" codec:"message,omitempty" mapstructure:"message,omitempty"`
}

type ErrorAPI interface {
	Error() string
	WithCause(err error) Error
}

// Error 包含业务状态的错误.
type Error struct {
	// StatusCode http status code
	StatusCode int
	// Code business service code
	Code int
	// Message error message
	Message string
	// Cause error
	Cause error
}

// Froze convert Error to ErrorAPI
func (e Error) Froze() ErrorAPI {
	return Error{
		StatusCode: e.StatusCode,
		Code:       e.Code,
		Message:    e.Message,
		Cause:      e.Cause,
	}
}

func (e Error) WithCause(err error) Error {
	if e.Cause == nil {
		return Error{
			StatusCode: e.StatusCode,
			Code:       e.Code,
			Message:    e.Message,
			Cause:      err,
		}
	}
	return Error{
		StatusCode: e.StatusCode,
		Code:       e.Code,
		Message:    e.Message,
		Cause:      errorx.Join(e.Cause, err),
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("gors.Error, StatusCode: %d, Code: %d, Message: %s", e.StatusCode, e.Code, e.Message)
}

func (e Error) Status() *Status {
	return &Status{Code: e.Code, Message: e.Message}
}

// Unwrap provides compatibility for Go 1.13 error chains.
func (e Error) Unwrap() error { return e.Cause }

// Is matches each error in the chain with the target value.
func (e Error) Is(err error) bool {
	ge := FromError(err)
	return ge.StatusCode == e.StatusCode && ge.Code == e.Code
}

// GRPCStatus returns the Status represented by gors.Error.
func (e Error) GRPCStatus() *gstatus.Status {
	s, _ := gstatus.New(status.ToGRPCCode(int(e.StatusCode)), e.Message).WithDetails(&status.Status{Code: int32(e.Code)})
	return s
}

// Format nolint: errcheck // WriteString could no check in pkg.
func (e Error) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%s\n", e.Error())
			if e.Cause != nil {
				fmt.Fprintf(s, "%+v", e.Cause)
			}
			return
		}
		if s.Flag('-') {
			fmt.Fprintf(s, "%s\n", e.Error())
			if e.Cause != nil {
				fmt.Fprintf(s, "%-v", e.Cause)
			}
			return
		}

		fallthrough
	case 's':
		_, _ = io.WriteString(s, e.Error())
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	}
}

// FromError 解析error为 Error.
// 解析 Error 或者 status.Status 为 Error, 其他类型的error返回 UnknownError.
func FromError(err error) Error {
	if err == nil {
		return Error{}
	}
	if v := errValue(); errors.As(err, &v) {
		return v
	}
	if v := new(Error); errors.As(err, &v) {
		return *v
	}
	ge, ok := gstatus.FromError(err)
	if !ok {
		return Error{
			StatusCode: UnknownStatusCode,
			Code:       UnknownCode,
			Message:    UnknownMessage,
			Cause:      err,
		}
	}
	ret := Error{StatusCode: status.FromGRPCCode(ge.Code()), Code: UnknownCode, Message: ge.Message()}
	for _, detail := range ge.Details() {
		switch d := detail.(type) {
		case *status.Status:
			ret.Code = int(d.Code)
		}
	}
	return ret
}

func ErrorFromMessage(msg string) (Error, bool) {
	if !msgRegExp.MatchString(msg) {
		return Error{}, false
	}
	subStrings := msgRegExp.FindAllStringSubmatch(msg, -1)
	if len(subStrings) != 1 {
		return Error{}, false
	}
	if len(subStrings[0]) != 4 {
		return Error{}, false
	}
	return Error{StatusCode: convx.ToInt(subStrings[0][1]), Code: convx.ToInt(subStrings[0][2]), Message: subStrings[0][3]}, true
}

func errValue() Error {
	return Error{}
}
