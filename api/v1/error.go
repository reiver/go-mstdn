package v1
type Error interface {
	error
	ErrHTTP() (httpStatusCode int)
	Unwrap() error
}
