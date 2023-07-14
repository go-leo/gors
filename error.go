package gors

import (
	"fmt"
	"github.com/go-leo/gox/convx"
	"regexp"
)

var msgRegExp = regexp.MustCompile("^gors.Error, Code: (\\d+), Message: (.+)$")

// Status 代表业务状态
type Status struct {
	Code    int    `json:"code,omitempty" yaml:"code,omitempty" xml:"code,omitempty" toml:"code,omitempty" codec:"code,omitempty" mapstructure:"code,omitempty"`
	Message string `json:"message,omitempty" yaml:"message,omitempty" xml:"message,omitempty" toml:"message,omitempty" codec:"message,omitempty" mapstructure:"message,omitempty"`
}

// Error 代表http状态码和业务状态
type Error struct {
	// StatusCode http status code
	StatusCode int
	// Code business service code
	Code int
	// Message error message
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("gors.Error, Code: %d, Message: %s", e.Code, e.Message)
}

func (e Error) Status() *Status {
	return &Status{Code: e.Code, Message: e.Message}
}

func ErrorFromMessage(msg string) (Error, bool) {
	if !msgRegExp.MatchString(msg) {
		return Error{}, false
	}
	subStrings := msgRegExp.FindAllStringSubmatch(msg, -1)
	if len(subStrings) != 1 {
		return Error{}, false
	}
	if len(subStrings[0]) != 3 {
		return Error{}, false
	}
	return Error{Code: convx.ToInt(subStrings[0][1]), Message: subStrings[0][2]}, true
}

func errorValue() Error { return Error{} }

func errorPointer() *Error { return &Error{} }
