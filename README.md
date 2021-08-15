# New Relic DB Stats Profiler

## Usage
```go
import (
    "time"

    "github.com/jmoiron/sqlx"
    dbprof "github.com/maruTA-bis5/nr-db-stats-profiler"
    "github.com/newrelic/go-agent/newrelic"
)

var db *sqlx.DB
var app *newrelic.Application

func Start() {
    dbprof.EnableDBStatsEventRecord(db, app, 10*time.Second)
}
```

## License
MIT


## Author
Takayuki Maruyama (@maruTA-bis5)
