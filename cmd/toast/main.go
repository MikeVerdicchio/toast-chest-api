package main

import (
	"database/sql"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/MikeVerdicchio/toast-chest-api/internal/controllers"
	"github.com/MikeVerdicchio/toast-chest-api/internal/health"
	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

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
		logger.Error("Could not setup database")
	}

	err = db.Ping()
	if err != nil {
		logger.Fatal("Could not establish a connection with the database")
	}

	// Setup router and handlers
	logger.Info("Setting up router and handlers")
	router := mux.NewRouter()
	baseHandler := controllers.NewBaseHandler(db, logger)
	configureHandlers(router, db, baseHandler)
	configureCORS(router)

	server := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("Listening on port " + port)
	log.Fatal(server.ListenAndServe())
}

// configureHandlers configures all routes for application
func configureHandlers(r *mux.Router, db *sql.DB, h *controllers.BaseHandler) {
	// Add health check endpoints
	healthHandler := health.ConfigureHealthHandler(db)
	r.HandleFunc("/live", healthHandler.LiveEndpoint).Methods(http.MethodGet)
	r.HandleFunc("/ready", healthHandler.ReadyEndpoint).Methods(http.MethodGet)

	// Random toast endpoint
	r.HandleFunc("/", h.RandomToastHandler).Methods(http.MethodGet)
}

// configureCORS configures CORS for application
func configureCORS(r *mux.Router) {
	uiURL := url.URL{
		Scheme: "https",
		Host:   os.Getenv("UI_HOST"),
	}
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{uiURL.String()}),
		handlers.AllowedMethods([]string{"OPTIONS", "GET"}),
	)

	r.Use(cors)
}
