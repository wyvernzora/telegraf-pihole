package scanner

import (
	"errors"
	"gotest.tools/v3/assert"
	"testing"
)

func TestIsNilValueError_NilValue(t *testing.T) {
	actual := isNilValueError(nil)
	assert.Equal(t, actual, false)
}

func TestIsNilValueError_GenericError(t *testing.T) {
	actual := isNilValueError(errors.New("test error"))
	assert.Equal(t, actual, false)
}

func TestIsNilValueError_NilValueError(t *testing.T) {
	actual := isNilValueError(nilValueError())
	assert.Equal(t, actual, true)
}

func TestUnexpectedValueError(t *testing.T) {
	err := unexpectedValueError(123)
	assert.ErrorType(t, &scanError{}, err)
	assert.Error(t, err, "unexpected value: 123")
}

func TestUnexpectedTypeError(t *testing.T) {
	err := unexpectedTypeError[int]("test")
	assert.ErrorType(t, &scanError{}, err)
	assert.Error(t, err, "expecting int, got: string")
}
