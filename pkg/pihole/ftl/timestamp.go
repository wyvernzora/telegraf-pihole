package ftl

import (
	"database/sql"
	"github.com/wyvernzora/telegraf-pihole/pkg/scanner"
	"time"
)

// ScanTimestamp returns a sql.Scanner instance that reads and populates the Query.Timestamp field.
//
// Converts the Unix epoch seconds value stored in the database into time.Time.
// Returns an error if the column type is not int64.
func ScanTimestamp(q *Query) sql.Scanner {
	return timestampScanner{&q.Timestamp}
}

type timestampScanner struct {
	ts *time.Time
}

func (ts timestampScanner) Scan(value interface{}) error {
	var v int64
	if err := scanner.New[int64]()(&v, value); err != nil {
		return err
	}
	*ts.ts = time.Unix(v, 0)
	return nil
}
