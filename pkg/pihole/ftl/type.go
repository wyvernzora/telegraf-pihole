package ftl

import (
	"database/sql"
	"github.com/wyvernzora/telegraf-pihole/pkg/scanner"
)

// Type of the incoming DNS request
type Type string

func (t Type) String() string {
	return string(t)
}

const (
	TypeA      Type = "A"
	TypeAAAA   Type = "AAAA"
	TypeANY    Type = "ANY"
	TypeSRV    Type = "SRV"
	TypeSOA    Type = "SOA"
	TypePTR    Type = "PTR"
	TypeTXT    Type = "TXT"
	TypeNAPTR  Type = "NAPTR"
	TypeMX     Type = "MX"
	TypeDS     Type = "DS"
	TypeRRSIG  Type = "RRSIG"
	TypeDNSKEY Type = "DNSKEY"
	TypeNS     Type = "NS"
	TypeOther  Type = "OTHER"
	TypeSVCB   Type = "SVCB"
	TypeHTTPS  Type = "HTTPS"
)

// ScanType returns a sql.Scanner instance that reads and populates the Query.Type field.
//
// Produces an error if the input value is null, column type is not int64, or if the value is not
// documented in https://docs.pi-hole.net/database/ftl/#supported-query-types
func ScanType(rec *Query) sql.Scanner {
	return typeScanner{v: &rec.Type}
}

var types = map[int64]Type{
	1:  TypeA,
	2:  TypeAAAA,
	3:  TypeANY,
	4:  TypeSRV,
	5:  TypeSOA,
	6:  TypePTR,
	7:  TypeTXT,
	8:  TypeNAPTR,
	9:  TypeMX,
	10: TypeDS,
	11: TypeRRSIG,
	12: TypeDNSKEY,
	13: TypeNS,
	14: TypeOther,
	15: TypeSVCB,
	16: TypeHTTPS,
}

type typeScanner struct {
	v *Type
}

func (s typeScanner) Scan(value interface{}) error {
	return scanner.WithMapping(types, scanner.New[int64]())(s.v, value)
}
