package scanner

import "fmt"

const (
	nilValueCode        = "nil_value"
	unexpectedValueCode = "unexpected_value"
	unexpectedTypeCode  = "unexpected_type"
)

type scanError struct {
	code    string
	message string
}

func (e *scanError) Code() string {
	return e.code
}

func (e *scanError) Error() string {
	return e.message
}

func nilValueError() error {
	return &scanError{
		code:    nilValueCode,
		message: "input value is nil",
	}
}

func isNilValueError(err error) bool {
	if nve, ok := err.(*scanError); ok {
		return nve.Code() == nilValueCode
	}
	return false
}

func unexpectedValueError(value interface{}) error {
	return &scanError{
		code:    unexpectedValueCode,
		message: fmt.Sprintf("unexpected value: %v", value),
	}
}

func unexpectedTypeError[T any](value interface{}) error {
	return &scanError{
		code:    unexpectedTypeCode,
		message: fmt.Sprintf("expecting %T, got: %T", *new(T), value),
	}
}
