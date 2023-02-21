package ftl

import (
	"database/sql"
	"github.com/wyvernzora/telegraf-pihole/pkg/scanner"
)

// ScanForward returns a sql.Scanner instance that reads and processes the Query.Timestamp field.
//
// Converts the epoch second timestamp into the proper nanosecond-based time.Time.
// `forward` column is nullable in the Pihole FTL database; when this scanner encounters a null value, it
// populates a special "NONE" value into the Query.Forward field.
// Returns an error if the column type is not a string.
func ScanForward(q *Query) sql.Scanner {
	return forwardScanner{&q.Forward}
}

// Internal struct to implement sql.Scanner for Query.Timestamp (type string)
type forwardScanner struct {
	fwd *string
}

func (fs forwardScanner) Scan(value interface{}) error {
	return scanner.WithDefault("none", scanner.New[string]())(fs.fwd, value)
}
