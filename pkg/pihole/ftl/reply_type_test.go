package ftl

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestReplyType_String(t *testing.T) {
	var actual = ReplyTypeCname.String()
	assert.Equal(t, actual, "CNAME")
}

func TestReplyType_Scan_NilValue(t *testing.T) {
	var query Query
	var scanner = ScanReplyType(&query)
	err := scanner.Scan(nil)

	assert.NilError(t, err)
	assert.Equal(t, query.ReplyType, ReplyTypeNone)
}

func TestReplyType_Scan_UnexpectedValue(t *testing.T) {
	var query Query
	var scanner = ScanReplyType(&query)
	err := scanner.Scan(int64(99))

	assert.ErrorContains(t, err, "unexpected value: 99")
}

func TestReplyType_Scan_UnexpectedType(t *testing.T) {
	var query Query
	var scanner = ScanReplyType(&query)
	err := scanner.Scan("test")

	assert.ErrorContains(t, err, "expecting int64, got: string")
}

func TestReplyType_Scan_ValidValue(t *testing.T) {
	var query Query
	var scanner = ScanReplyType(&query)
	err := scanner.Scan(int64(3))

	assert.NilError(t, err)
	assert.Equal(t, query.ReplyType, ReplyTypeCname)
}
