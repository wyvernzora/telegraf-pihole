package ftl

import (
	"database/sql"
	"github.com/wyvernzora/telegraf-pihole/pkg/scanner"
)

// ReplyType that Pihole responded to the incoming DNS ftl with
type ReplyType string

func (rt ReplyType) String() string {
	return string(rt)
}

const (
	ReplyTypeUnknown  ReplyType = "UNKNOWN"
	ReplyTypeNoData   ReplyType = "NODATA"
	ReplyTypeNxDomain ReplyType = "NXDOMAIN"
	ReplyTypeCname    ReplyType = "CNAME"
	ReplyTypeIP       ReplyType = "IP"
	ReplyTypeDomain   ReplyType = "DOMAIN"
	ReplyTypeRRName   ReplyType = "RRNAME"
	ReplyTypeServFail ReplyType = "SERVFAIL"
	ReplyTypeRefused  ReplyType = "REFUSED"
	ReplyTypeNotImp   ReplyType = "NOTIMP"
	ReplyTypeOther    ReplyType = "OTHER"
	ReplyTypeDnsSec   ReplyType = "DNSSEC"
	ReplyTypeNone     ReplyType = "NONE"
	ReplyTypeBlob     ReplyType = "BLOB"
)

// ScanReplyType returns a sql.Scanner instance that reads and populates the Query.ReplyType field.
//
// `reply_type` column is nullable in the Pihole FTL database; when this scanner encounters a null value,
// it populates a special "NONE" value into the Query.ReplyType field.
// Returns an error if the column type is not int64, or if the value is not documented in
// https://docs.pi-hole.net/database/ftl/#supported-reply-types
func ScanReplyType(rec *Query) sql.Scanner {
	return replyTypeScanner{v: &rec.ReplyType}
}

var replyTypes = map[int64]ReplyType{
	0:  ReplyTypeUnknown,
	1:  ReplyTypeNoData,
	2:  ReplyTypeNxDomain,
	3:  ReplyTypeCname,
	4:  ReplyTypeIP,
	5:  ReplyTypeDomain,
	6:  ReplyTypeRRName,
	7:  ReplyTypeServFail,
	8:  ReplyTypeRefused,
	9:  ReplyTypeNotImp,
	10: ReplyTypeOther,
	11: ReplyTypeDnsSec,
	12: ReplyTypeNone,
	13: ReplyTypeBlob,
}

type replyTypeScanner struct {
	v *ReplyType
}

// Scan implements the sql.Scanner interface for *ReplyType
func (s replyTypeScanner) Scan(value interface{}) error {
	return scanner.WithDefault(ReplyTypeNone, scanner.WithMapping(replyTypes, scanner.New[int64]()))(s.v, value)
}
