package controllers

import (
	"database/sql"

	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	// Used for defining sql.DB
	_ "github.com/lib/pq"
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
	if err := h.db.Ping(); err != nil {
		h.logger.Error("Could not ping database")
	}

	h.logger.Info("RandomToastHandler handled")
	fmt.Fprint(w, DefaultToast)
}
