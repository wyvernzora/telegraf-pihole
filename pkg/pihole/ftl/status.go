package ftl

import (
	"database/sql"
	"github.com/wyvernzora/telegraf-pihole/pkg/scanner"
)

// Decision that Pihole made regarding the ftl in question
// For more details, see https://docs.pi-hole.net/database/ftl/#supported-status-types
type Decision string

func (d Decision) String() string {
	return string(d)
}

// Reason that Pihole made the Decision in question
// For more details, see https://docs.pi-hole.net/database/ftl/#supported-status-types
type Reason string

func (r Reason) String() string {
	return string(r)
}

const (
	DecisionAllowed Decision = "Allowed"
	DecisionBlocked Decision = "Blocked"
	DecisionUnknown Decision = "Unknown"
)

const (
	ReasonUnknown          Reason = "Unknown"
	ReasonGravity          Reason = "Gravity"
	ReasonForwarded        Reason = "Forwarded"
	ReasonCached           Reason = "Cached"
	ReasonRegexMatch       Reason = "RegexMatch"
	ReasonExactMatch       Reason = "ExactMatch"
	ReasonByUpstream       Reason = "ByUpstream"
	ReasonZeroIP           Reason = "ZeroIP"
	ReasonNxDomain         Reason = "NXDOMAIN"
	ReasonCnameGravity     Reason = "CnameGravity"
	ReasonCnameRegexMatch  Reason = "CnameRegexMatch"
	ReasonCnameExactMatch  Reason = "CnameExactMatch"
	ReasonRetried          Reason = "Retried"
	ReasonRetriedIgnored   Reason = "RetriedIgnored"
	ReasonAlreadyForwarded Reason = "AlreadyForwarded"
	ReasonDatabaseBusy     Reason = "DatabaseBusy"
	ReasonSpecialDomain    Reason = "SpecialDomain"
	ReasonCachedStale      Reason = "CachedStale"
)

// ScanStatus returns a sql.Scanner instance that reads and populates the Query.Decision
// and Query.Reason fields.
//
// Produces an error if encounters a null value, if column type is not int64, or if the value
// is not documented by Pihole.
func ScanStatus(rec *Query) sql.Scanner {
	return statusScanner{d: &rec.Decision, r: &rec.Reason}
}

type status struct {
	Decision
	Reason
}

var statuses = map[int64]status{
	0:  {DecisionUnknown, ReasonUnknown},
	1:  {DecisionBlocked, ReasonGravity},
	2:  {DecisionAllowed, ReasonForwarded},
	3:  {DecisionAllowed, ReasonCached},
	4:  {DecisionBlocked, ReasonRegexMatch},
	5:  {DecisionBlocked, ReasonExactMatch},
	6:  {DecisionBlocked, ReasonByUpstream},
	7:  {DecisionBlocked, ReasonZeroIP},
	8:  {DecisionBlocked, ReasonNxDomain},
	9:  {DecisionBlocked, ReasonCnameGravity},
	10: {DecisionBlocked, ReasonCnameRegexMatch},
	11: {DecisionBlocked, ReasonCnameExactMatch},
	12: {DecisionAllowed, ReasonRetried},
	13: {DecisionAllowed, ReasonRetriedIgnored},
	14: {DecisionAllowed, ReasonAlreadyForwarded},
	15: {DecisionBlocked, ReasonDatabaseBusy},
	16: {DecisionBlocked, ReasonSpecialDomain},
	17: {DecisionAllowed, ReasonCachedStale},
}

// Internal struct to implement sql.Scanner for Query.Decision and Query.Reason
type statusScanner struct {
	d *Decision
	r *Reason
}

func (s statusScanner) Scan(value interface{}) error {
	var status status
	if err := scanner.WithMapping(statuses, scanner.New[int64]())(&status, value); err != nil {
		return err
	}
	*s.d = status.Decision
	*s.r = status.Reason
	return nil
}
