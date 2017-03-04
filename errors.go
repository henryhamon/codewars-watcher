package main

type httpError struct {
	error
	statusCode int
}

// NewHTTPError API error
func NewHTTPError(err error, code int) error {
	return httpError{err, code}
}

type decodingError struct {
	error
	objectType string
}

// NewDecodingError error used when
// get error decoding objects
func NewDecodingError(err error, objectType string) error {
	return decodingError{err, objectType}
}
