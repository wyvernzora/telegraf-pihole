package scanner

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestScanner_NilValue(t *testing.T) {
	var output string
	err := New[string]()(&output, nil)

	assert.Equal(t, isNilValueError(err), true)
}

func TestScanner_UnexpectedType(t *testing.T) {
	var output string
	err := New[string]()(&output, 42)

	assert.Error(t, err, "expecting string, got: int")
}

func TestScanner_ValidValue(t *testing.T) {
	var output string
	err := New[string]()(&output, "foo")

	assert.NilError(t, err)
	assert.Equal(t, output, "foo")
}

func TestWithDefault_NilValue(t *testing.T) {
	var output string
	err := WithDefault("bar", New[string]())(&output, nil)

	assert.NilError(t, err)
	assert.Equal(t, output, "bar")
}

func TestWithDefault_UnexpectedType(t *testing.T) {
	var output string
	err := WithDefault("bar", New[string]())(&output, 42)

	assert.Error(t, err, "expecting string, got: int")
}

var testMap = map[int]string{
	0: "zero",
	1: "one",
}

func TestWithMapping_ValueInMap(t *testing.T) {
	var output string
	err := WithMapping(testMap, New[int]())(&output, 0)

	assert.NilError(t, err)
	assert.Equal(t, output, "zero")
}

func TestWithMapping_UnexpectedValue(t *testing.T) {
	var output string
	err := WithMapping(testMap, New[int]())(&output, 9)

	assert.Error(t, err, "unexpected value: 9")
}

func TestWithMapping_NilValue(t *testing.T) {
	var output string
	err := WithMapping(testMap, New[int]())(&output, nil)

	assert.Equal(t, isNilValueError(err), true)
}
