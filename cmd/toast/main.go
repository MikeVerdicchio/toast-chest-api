package main

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/MikeVerdicchio/toast-chest-api/internal/controllers"
	"github.com/MikeVerdicchio/toast-chest-api/internal/health"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	// Used for defining sql.DB
	_ "github.com/lib/pq"
)

const (
	// ApplicationName is the name of the service
	ApplicationName = "toast-api"
)

func main() {
	logger := log.WithFields(log.Fields{"app": ApplicationName})

	// Read environment variables - TODO use viper
	connStr := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")

	// Setup database connection
	logger.Info("Setting up database connection")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("Could not connect to database")
	}

	// Setup router and handlers
	logger.Info("Setting up router and handlers")
	router := mux.NewRouter()
	baseHandler := controllers.NewBaseHandler(db, logger)
	ConfigureHandlers(router, db, baseHandler)

	server := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("Listening on port " + port)
	log.Fatal(server.ListenAndServe())
}

// ConfigureHandlers configures all routes for application
func ConfigureHandlers(r *mux.Router, db *sql.DB, h *controllers.BaseHandler) {
	// Add health check endpoints
	healthHandler := health.ConfigureHealthHandler(db)
	r.HandleFunc("/live", healthHandler.LiveEndpoint).Methods(http.MethodGet)
	r.HandleFunc("/ready", healthHandler.ReadyEndpoint).Methods(http.MethodGet)

	// Random toast endpoint
	r.HandleFunc("/", h.RandomToastHandler).Methods(http.MethodGet)
}
