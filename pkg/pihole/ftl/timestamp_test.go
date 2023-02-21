package ftl

import (
	"gotest.tools/v3/assert"
	"testing"
	"time"
)

func TestTimestamp_Scan_NilValue(t *testing.T) {
	var query Query
	var scanner = ScanTimestamp(&query)
	err := scanner.Scan(nil)

	assert.ErrorContains(t, err, "input value is nil")
}

func TestTimestamp_Scan_UnexpectedType(t *testing.T) {
	var query Query
	var scanner = ScanTimestamp(&query)
	err := scanner.Scan("test")

	assert.ErrorContains(t, err, "expecting int64, got: string")
}

func TestTimestamp_Scan_ValidValue(t *testing.T) {
	ts := time.Unix(1645597342, 0)

	var query Query
	var scanner = ScanTimestamp(&query)
	err := scanner.Scan(ts.Unix())

	assert.NilError(t, err)
	assert.Equal(t, query.Timestamp, ts)
}
