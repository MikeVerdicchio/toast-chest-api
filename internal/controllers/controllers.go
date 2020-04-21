package controllers

import (
	"database/sql"
	"encoding/json"

	"net/http"

	"github.com/MikeVerdicchio/toast-chest-api/internal/toast"
	log "github.com/sirupsen/logrus"
)

const (
	// DefaultToast is the default toast to show when there is an error
	DefaultToast = "Here's to friendship, here's to great times, and here's to when the Toast Chest isn't around."
)

// BaseHandler will hold db connection and other configurations
type BaseHandler struct {
	db     *sql.DB
	logger *log.Entry
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(db *sql.DB, logger *log.Entry) *BaseHandler {
	return &BaseHandler{
		db:     db,
		logger: logger,
	}
}

// RandomToastHandler returns a random toast
func (h *BaseHandler) RandomToastHandler(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling RandomToastHandler request")

	var t toast.Toast
	if err := t.GetRandomToast(h.db); err != nil {
		h.logger.Error("Could not pull toast from database")
		t.Toast = DefaultToast
	}

	returnJSON(w, http.StatusOK, t)
}

func returnJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		log.WithError(err).Error("Error writing response")
	}
}

func returnJSONError(w http.ResponseWriter, code int, message string) {
	returnJSON(w, code, map[string]string{"error": message})
}
