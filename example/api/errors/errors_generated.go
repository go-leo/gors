// Code generated by "gors-error -type=int  "; DO NOT EDIT.
package errors

import "github.com/go-leo/gors"

var (
	ErrUnknown = gors.Error{
		StatusCode: 500,
		Code:       Unknown,
		Message:    "Internal server error",
	}.Froze()

	ErrBind = gors.Error{
		StatusCode: 400,
		Code:       Bind,
		Message:    "Error occurred while binding the request body to the struct",
	}.Froze()

	ErrValidation = gors.Error{
		StatusCode: 400,
		Code:       Validation,
		Message:    "Validation failed",
	}.Froze()

	ErrAccountAuthTypeInvalid = gors.Error{
		StatusCode: 400,
		Code:       AccountAuthTypeInvalid,
		Message:    "Account AuthType not support",
	}.Froze()

	ErrUserNotFound = gors.Error{
		StatusCode: 400,
		Code:       UserNotFound,
		Message:    "User Not Found",
	}.Froze()

	ErrUserDisabled = gors.Error{
		StatusCode: 400,
		Code:       UserDisabled,
		Message:    "User disabled",
	}.Froze()
)
