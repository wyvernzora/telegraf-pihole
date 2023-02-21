package ftl

import (
	"gotest.tools/v3/assert"
	"testing"
	"time"
)

func TestReplyTime_Scan_NilValue(t *testing.T) {
	var query Query
	var scanner = ScanReplyTime(&query)
	err := scanner.Scan(nil)

	assert.NilError(t, err)
	assert.Equal(t, query.ReplyTime, time.Duration(0))
}

func TestReplyTime_Scan_UnexpectedType(t *testing.T) {
	var query Query
	var scanner = ScanReplyTime(&query)
	err := scanner.Scan("test")

	assert.ErrorContains(t, err, "expecting float64, got: string")
}

func TestReplyTime_Scan_ValidValue(t *testing.T) {
	var query Query
	var scanner = ScanReplyTime(&query)
	err := scanner.Scan(100.5)

	assert.NilError(t, err)
	assert.Equal(t, query.ReplyTime, 100500*time.Millisecond)
}
