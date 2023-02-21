package ftl

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestForward_Scan_NilValue(t *testing.T) {
	var query Query
	var scanner = ScanForward(&query)
	err := scanner.Scan(nil)

	assert.NilError(t, err)
	assert.Equal(t, query.Forward, "none")
}

func TestForward_Scan_UnexpectedType(t *testing.T) {
	var query Query
	var scanner = ScanForward(&query)
	err := scanner.Scan(1)

	assert.ErrorContains(t, err, "expecting string, got: int")
}

func TestForward_Scan_ValidValue(t *testing.T) {
	var query Query
	var scanner = ScanForward(&query)
	err := scanner.Scan("1.2.3.4")

	assert.NilError(t, err)
	assert.Equal(t, query.Forward, "1.2.3.4")
}
