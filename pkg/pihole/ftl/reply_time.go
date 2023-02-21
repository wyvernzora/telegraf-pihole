package ftl

import (
	"database/sql"
	"github.com/wyvernzora/telegraf-pihole/pkg/scanner"
	"time"
)

// ScanReplyTime returns a sql.Scanner instance that reads and populates the Query.ReplyTime field.
//
// Converts the float64 seconds value stored in the database into time.Duration.
// `reply_time` column is nullable in the Pihole FTL database; when this scanner encounters a null value, it
// populates a zero duration into the Query.ReplyTime field.
// Returns an error if the column type is not float64.
func ScanReplyTime(q *Query) sql.Scanner {
	return replyTimeScanner{&q.ReplyTime}
}

type replyTimeScanner struct {
	rt *time.Duration
}

func (rts replyTimeScanner) Scan(value interface{}) error {
	var t float64
	if err := scanner.WithDefault(0, scanner.New[float64]())(&t, value); err != nil {
		return err
	}
	*rts.rt = time.Duration(t * float64(time.Second))
	return nil
}
