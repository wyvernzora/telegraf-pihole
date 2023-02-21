package ftl

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestDecision_String(t *testing.T) {
	var actual = DecisionAllowed.String()
	assert.Equal(t, actual, "Allowed")
}

func TestReason_String(t *testing.T) {
	var actual = ReasonForwarded.String()
	assert.Equal(t, actual, "Forwarded")
}

func TestStatus_Scan_NilValue(t *testing.T) {
	var query Query
	var scanner = ScanStatus(&query)
	err := scanner.Scan(nil)

	assert.ErrorContains(t, err, "input value is nil")
}

func TestStatus_Scan_UnexpectedValue(t *testing.T) {
	var query Query
	var scanner = ScanStatus(&query)
	err := scanner.Scan(int64(99))

	assert.ErrorContains(t, err, "unexpected value: 99")
}

func TestStatus_Scan_UnexpectedType(t *testing.T) {
	var query Query
	var scanner = ScanStatus(&query)
	err := scanner.Scan("test")

	assert.ErrorContains(t, err, "expecting int64, got: string")
}

func TestStatus_Scan_ValidValue(t *testing.T) {
	var query Query
	var scanner = ScanStatus(&query)
	err := scanner.Scan(int64(2))

	assert.NilError(t, err)
	assert.Equal(t, query.Decision, DecisionAllowed)
	assert.Equal(t, query.Reason, ReasonForwarded)
}
