package errors

//go:generate gors-error -type=int

// base: base errors.
const (
	// Unknown - 500: Internal server error.
	Unknown int = iota + 100001

	// Bind - 400: Error occurred while binding the request body to the struct.
	Bind

	// Validation - 400: Validation failed.
	Validation
)

// Account-server: Account errors.
const (
	// AccountAuthTypeInvalid - 400: Account AuthType not support.
	AccountAuthTypeInvalid int = iota + 110001

	// UserNotFound - 400: User Not Found.
	UserNotFound

	// UserDisabled - 400: User disabled.
	UserDisabled
)
