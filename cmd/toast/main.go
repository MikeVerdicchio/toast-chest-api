package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/MikeVerdicchio/toast-chest-api/internal/health"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const (
	// ApplicationName is the name of the service
	ApplicationName = "toast-api"
	// DefaultToast is the default toast to show when there is an error
	DefaultToast = "Here's to friendship, here's to great times, and here's to when the Toast Chest isn't around."
)

func main() {
	logger := log.WithFields(log.Fields{"app": ApplicationName})
	logger.Info("Loading application")

	router := mux.NewRouter()
	ConfigureHandlers(router)

	port := os.Getenv("LISTEN_PORT")
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
func ConfigureHandlers(r *mux.Router) {
	// Add health check endpoints
	healthHandler := health.ConfigureHealthHandler()
	r.HandleFunc("/live", healthHandler.LiveEndpoint).Methods(http.MethodGet)
	r.HandleFunc("/ready", healthHandler.ReadyEndpoint).Methods(http.MethodGet)

	// Random toast endpoint
	r.HandleFunc("/", RandomToastHandler).Methods(http.MethodGet)
}

// RandomToastHandler returns a random toast
func RandomToastHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, DefaultToast)
}
