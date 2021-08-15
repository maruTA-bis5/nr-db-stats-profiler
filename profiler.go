package dbprof

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func EnableDBStatsEventRecord(db *sqlx.DB, app *newrelic.Application, period time.Duration) {
	go func() {
		for {
			stats := db.Stats()
			app.RecordCustomEvent("DBStats", map[string]interface{} {
				"MaxOpenConnection": stats.MaxOpenConnections,
				"Idle":              stats.Idle,
				"OpenConnections":   stats.OpenConnections,
				"InUse":             stats.InUse,
			})
			time.Sleep(period)
		}
	}()
}
