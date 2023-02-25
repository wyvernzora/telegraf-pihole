package pihole

import (
	"github.com/influxdata/telegraf"
	"github.com/wyvernzora/telegraf-pihole/pkg/pihole/ftl"
	"path"
	"time"
)

func (p *Pihole) gatherQueryLogs(a telegraf.Accumulator) (err error) {
	if !p.GatherQueryLogs {
		return nil
	}

	db, err := p.openDatabase(path.Join(p.PiholeConfigRoot, "pihole-FTL.db"))
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
		writeQueryLogEntry(a, buffer[i])
	}
	p.position = reader.Position()
	p.Log.Infof("Processed %d queries; reader position is at %d", n, reader.Position())

	if p.GatherPluginMetrics {
		var lag = 0 * time.Millisecond
		if n > 0 {
			lag = time.Now().Sub(buffer[n-1].Timestamp)
		}
		a.AddFields("pihole_plugin", map[string]interface{}{
			"lag":   lag.Milliseconds(),
			"count": n,
		}, map[string]string{})
	}

	return nil
}

func writeQueryLogEntry(a telegraf.Accumulator, query ftl.Query) {
	a.AddFields("query_log", map[string]interface{}{
		"reply_time": query.ReplyTime.Milliseconds(),
	}, map[string]string{
		"type":       query.Type.String(),
		"decision":   query.Decision.String(),
		"reason":     query.Reason.String(),
		"domain":     query.Domain,
		"client":     query.Client,
		"forward":    query.Forward,
		"reply_type": query.ReplyType.String(),
		"dnssec":     query.DnsSecStatus.String(),
	}, enhanceQueryTimestamp(query))
}

// enhanceQueryTimestamp artificially increases timestamp resolution by combining it with ID.
// Pihole FTL stores ftl timestamp using UNIX seconds, but second-level resolution
// is not sufficient for metrics systems such as InfluxDB. Therefore, using the record
// ID to artificially assign the millisecond component of the timestamp.
func enhanceQueryTimestamp(record ftl.Query) time.Time {
	var timestamp64 = record.Timestamp.UnixMilli() + record.Id%1000
	return time.UnixMilli(timestamp64)
}
