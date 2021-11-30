package errors

/*
	Custom message made in case of undefined error.
*/
const DEFAULT_MESSAGE = "SOMETHING WENT TERRIBLY WRONG"

/*
	Custom error made for wrapping CSV handling errors.
*/
type CSVFormatError struct {
	Message string
	Err     error
}

/*
	Custom error made in case the searched item is not found.
*/
type EntityNotFoundError struct {
	Message string
	Err     error
}

/*
	Custom error that wraps the errors made on request.
*/
type InvalidRequestError struct {
	Message string
	Err     error
}

/*
	Custom error made in case any business rule is not met.
*/
type DomainValidationError struct {
	Message string
	Err     error
}

// Returns custom message.
func (err CSVFormatError) Error() string {
	return err.Message
}

// Returns original error.
func (err CSVFormatError) Unwrap() error {
	return err.Err
}

// Returns custom message.
func (err EntityNotFoundError) Error() string {
	return err.Message
}

// Returns original error.
func (err EntityNotFoundError) Unwrap() error {
	return err.Err
}

// Returns custom message.
func (err InvalidRequestError) Error() string {
	return err.Message
}

// Returns original error.
func (err InvalidRequestError) Unwrap() error {
	return err.Err
}

// Returns custom message.
func (err DomainValidationError) Error() string {
	return err.Message
}

// Returns original error.
func (err DomainValidationError) Unwrap() error {
	return err.Err
}
