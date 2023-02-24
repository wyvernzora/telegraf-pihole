package pihole

import (
	"database/sql"
	"github.com/influxdata/telegraf"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wyvernzora/telegraf-pihole/pkg/pihole/ftl"
	"path"
	"time"
)

func (p *Pihole) gatherQueryLogs(a telegraf.Accumulator) (err error) {
	if !p.GatherQueryLogs {
		return nil
	}

	db, err := sql.Open("sqlite3", path.Join(p.PiholeConfigRoot, "pihole-FTL.db")+"?mode=ro")
	if err != nil {
		return
	}
	defer db.Close()

	reader := ftl.NewReader(db, p.position)
	var buffer = make([]ftl.Query, p.BatchSizeLimit)
	n, err := reader.Read(buffer)
	if err != nil {
		return
	}

	for i := 0; i < n; i++ {
		query := buffer[i]
		a.AddFields("query_log", map[string]interface{}{
			"reply_time": query.ReplyTime.Milliseconds(),
		}, map[string]string{
			"type":       query.Type.String(),
			"decision":   query.Decision.String(),
			"domain":     query.Domain,
			"client":     query.Client,
			"forward":    query.Forward,
			"reply_type": query.ReplyType.String(),
			"dnssec":     query.DnsSecStatus.String(),
		}, enhanceQueryTimestamp(query))
	}
	p.position = reader.Position()
	p.Log.Infof("Processed %d queries; reader position is at %d", n, reader.Position())

	return nil
}

// enhanceQueryTimestamp artificially increases timestamp resolution by combining it with ID.
// Pihole FTL stores ftl timestamp using UNIX seconds, but second-level resolution
// is not sufficient for metrics systems such as InfluxDB. Therefore, using the record
// ID to artificially assign the millisecond component of the timestamp.
func enhanceQueryTimestamp(record ftl.Query) time.Time {
	var timestamp64 = record.Timestamp.UnixMilli() + record.Id%1000
	return time.UnixMilli(timestamp64)
}
