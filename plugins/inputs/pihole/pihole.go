package pihole

import (
	"database/sql"
	_ "embed"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/filter"
	"github.com/influxdata/telegraf/plugins/inputs"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed pihole.toml
var sampleConfig string

func init() {
	inputs.Add("pihole", func() telegraf.Input { return New() })
}

func New() *Pihole {
	return &Pihole{
		PiholeConfigRoot:    "/etc/pihole",
		BatchSizeLimit:      5000,
		GatherQueryLogs:     true,
		GatherAdListMetrics: true,
		GatherPluginMetrics: true,

		// This is the only file IO bit that is difficult to mock
		// Putting it here to make testing the rest of the code easier
		openDatabase: func(filepath string) (*sql.DB, error) {
			return sql.Open("sqlite3", filepath+"?mode=ro")
		},
	}
}

type Pihole struct {
	PiholeConfigRoot    string   `toml:"pihole_config_root"`
	BatchSizeLimit      int      `toml:"batch_size_limit"`
	GatherQueryLogs     bool     `toml:"gather_query_logs"`
	GatherAdListMetrics bool     `toml:"gather_adlist_metrics"`
	GatherPluginMetrics bool     `toml:"gather_plugin_metrics"`
	IncludeTags         []string `toml:"include_tags"`
	ExcludeTags         []string `toml:"exclude_tags"`

	Log telegraf.Logger `toml:"-"`

	position  int64
	tagFilter filter.Filter

	openDatabase func(string) (*sql.DB, error)
}

func (p *Pihole) SampleConfig() string {
	return sampleConfig
}

func (p *Pihole) Description() string {
	return "Pull Pihole query logs and stats"
}

func (p *Pihole) Gather(a telegraf.Accumulator) (err error) {
	p.Log.Info("Starting metrics gathering cycle")
	if err = p.gatherQueryLogs(a); err != nil {
		return
	}
	return
}
