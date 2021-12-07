package errors

// Custom message made in case of undefined error.
const DefaultMessage = "Something went terribly wrong."

// Custom error made for wrapping CSV handling errors.
type ErrCSVFormat struct {
	Message string
	Err     error
}

// Custom error made for wrapping HTTP request errors.
type ErrHTTPRequest struct {
	Message string
	Err     error
}

// Custom error made in case the searched item is not found.
type ErrEntityNotFound struct {
	Message string
	Err     error
}

// Custom error that wraps the errors made on request.
type ErrInvalidRequest struct {
	Message string
	Err     error
}

// Custom error made in case any business rule is not met.
type ErrDomainValidation struct {
	Message string
	Err     error
}

// Custom error meant to be a wraper of an error from a repository.
type ErrRepositoryWrapper struct {
	Message string
	Err     error
}

// Custom error meant to be a wraper of an error from a infrastructure datastore.
type ErrDatastoreWrapper struct {
	Message string
	Err     error
}

// Custom error made to wrap system config errors.
type ErrSystemConfig struct {
	Message string
	Err     error
}

// Returns custom message.
func (err ErrCSVFormat) Error() string {
	return err.Message
}

// Returns original error.
func (err ErrCSVFormat) Unwrap() error {
	return err.Err
}

// Returns custom message.
func (err ErrHTTPRequest) Error() string {
	return err.Message
}

// Returns original error.
func (err ErrHTTPRequest) Unwrap() error {
	return err.Err
}

// Returns custom message.
func (err ErrEntityNotFound) Error() string {
	return err.Message
}

// Returns original error.
func (err ErrEntityNotFound) Unwrap() error {
	return err.Err
}

// Returns custom message.
func (err ErrInvalidRequest) Error() string {
	return err.Message
}

// Returns original error.
func (err ErrInvalidRequest) Unwrap() error {
	return err.Err
}

// Returns custom message.
func (err ErrDomainValidation) Error() string {
	return err.Message
}

// Returns original error.
func (err ErrDomainValidation) Unwrap() error {
	return err.Err
}

// Returns custom message.
func (err ErrRepositoryWrapper) Error() string {
	return err.Message
}

// Returns original error.
func (err ErrRepositoryWrapper) Unwrap() error {
	return err.Err
}

// Returns custom message.
func (err ErrDatastoreWrapper) Error() string {
	return err.Message
}

// Returns original error.
func (err ErrDatastoreWrapper) Unwrap() error {
	return err.Err
}

// Returns custom message.
func (err ErrSystemConfig) Error() string {
	return err.Message
}

// Returns original error.
func (err ErrSystemConfig) Unwrap() error {
	return err.Err
}
