package health

import (
	"github.com/heptiolabs/healthcheck"
)

// ConfigureHealthHandler creates and configures a custom healthcheck handler
func ConfigureHealthHandler() healthcheck.Handler {
	health := healthcheck.NewHandler()

	// TODO: add database health ping

	return health
}
