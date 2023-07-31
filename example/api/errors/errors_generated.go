// Code generated by "gors-error -type=int  "; DO NOT EDIT.
package errors

import "github.com/go-leo/gors"

var (
	ErrUnknown = gors.Error{
		StatusCode: 500,
		Code:       Unknown,
		Message:    "Internal server error",
	}

	ErrBind = gors.Error{
		StatusCode: 400,
		Code:       Bind,
		Message:    "Error occurred while binding the request body to the struct",
	}

	ErrValidation = gors.Error{
		StatusCode: 400,
		Code:       Validation,
		Message:    "Validation failed",
	}

	ErrAccountAuthTypeInvalid = gors.Error{
		StatusCode: 400,
		Code:       AccountAuthTypeInvalid,
		Message:    "Account AuthType not support",
	}

	ErrUserNotFound = gors.Error{
		StatusCode: 400,
		Code:       UserNotFound,
		Message:    "User Not Found",
	}

	ErrUserDisabled = gors.Error{
		StatusCode: 400,
		Code:       UserDisabled,
		Message:    "User disabled",
	}
)