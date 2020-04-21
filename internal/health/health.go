package health

import (
	"database/sql"
	"time"

	"github.com/heptiolabs/healthcheck"

	// Used for defining sql.DB
	_ "github.com/lib/pq"
)

// ConfigureHealthHandler creates and configures a custom healthcheck handler
func ConfigureHealthHandler(db *sql.DB) healthcheck.Handler {
	health := healthcheck.NewHandler()
	health.AddReadinessCheck("database", healthcheck.DatabasePingCheck(db, 1*time.Second))
	return health
}
