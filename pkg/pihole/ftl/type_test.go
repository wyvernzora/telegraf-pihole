package ftl

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestType_Scan_NilValue(t *testing.T) {
	var query Query
	var scanner = ScanType(&query)
	err := scanner.Scan(nil)

	assert.ErrorContains(t, err, "input value is nil")
}

func TestType_Scan_UnexpectedValue(t *testing.T) {
	var query Query
	var scanner = ScanType(&query)
	err := scanner.Scan(int64(99))

	assert.ErrorContains(t, err, "unexpected value: 99")
}

func TestType_Scan_UnexpectedType(t *testing.T) {
	var query Query
	var scanner = ScanType(&query)
	err := scanner.Scan("test")

	assert.ErrorContains(t, err, "expecting int64, got: string")
}

func TestType_Scan_ValidValue(t *testing.T) {
	var query Query
	var scanner = ScanType(&query)
	err := scanner.Scan(int64(2))

	assert.NilError(t, err)
	assert.Equal(t, query.Type, TypeAAAA)
}

func TestType_String(t *testing.T) {
	actual := TypeSVCB.String()
	assert.Equal(t, actual, "SVCB")
}
