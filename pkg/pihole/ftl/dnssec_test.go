package ftl

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestDnsSec_String(t *testing.T) {
	actual := DnsSecSecure.String()
	assert.Equal(t, actual, "Secure")
}

func TestDnsSec_Scan_NilValue(t *testing.T) {
	var query Query
	var scanner = ScanDnsSec(&query)
	err := scanner.Scan(nil)

	assert.NilError(t, err)
	assert.Equal(t, query.DnsSecStatus, DnsSecNone)
}

func TestDnsSec_Scan_UnexpectedValue(t *testing.T) {
	var query Query
	var scanner = ScanDnsSec(&query)
	err := scanner.Scan(int64(99))

	assert.ErrorContains(t, err, "unexpected value: 99")
}

func TestDnsSec_Scan_UnexpectedType(t *testing.T) {
	var query Query
	var scanner = ScanDnsSec(&query)
	err := scanner.Scan("test")

	assert.ErrorContains(t, err, "expecting int64, got: string")
}

func TestDnsSec_Scan_ValidValue(t *testing.T) {
	var query Query
	var scanner = ScanDnsSec(&query)
	err := scanner.Scan(int64(1))

	assert.NilError(t, err)
	assert.Equal(t, query.DnsSecStatus, DnsSecSecure)
}
