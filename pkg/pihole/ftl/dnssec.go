package ftl

import (
	"database/sql"
	"github.com/wyvernzora/telegraf-pihole/pkg/scanner"
)

// DnsSecStatus represents the state of DNSSEC for the given request
type DnsSecStatus string

func (dss DnsSecStatus) String() string {
	return string(dss)
}

const (
	// DnsSecNone denotes the situation where no DnsSecStatus is available
	DnsSecNone      DnsSecStatus = "None"
	DnsSecUnknown   DnsSecStatus = "Unknown"
	DnsSecSecure    DnsSecStatus = "Secure"
	DnsSecInsecure  DnsSecStatus = "Insecure"
	DnsSecBogus     DnsSecStatus = "Bogus"
	DnsSecAbandoned DnsSecStatus = "Abandoned"
)

// ScanDnsSec returns a sql.Scanner instance that reads and populates the Query.DnsSecStatus field.
//
// `dnssec` column is nullable in the Pihole FTL database; when this scanner encounters a null value, it
// populates a special "NONE" value into the Query.DnsSecStatus field.
// Returns an error if the column type is not int64, or if the value is not documented in
// https://docs.pi-hole.net/database/ftl/#dnssec-status
func ScanDnsSec(rec *Query) sql.Scanner {
	return dnsSecScanner{v: &rec.DnsSecStatus}
}

var dnsSecStatuses = map[int64]DnsSecStatus{
	0: DnsSecUnknown,
	1: DnsSecSecure,
	2: DnsSecInsecure,
	3: DnsSecBogus,
	4: DnsSecAbandoned,
}

type dnsSecScanner struct {
	v *DnsSecStatus
}

func (s dnsSecScanner) Scan(value interface{}) error {
	return scanner.WithDefault(DnsSecNone, scanner.WithMapping(dnsSecStatuses, scanner.New[int64]()))(s.v, value)
}
