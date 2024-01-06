package rest

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Hanlder defines handler
type HealthHandler struct {
}

// NewHandler ...
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Router ...
func (h *HealthHandler) Router() http.Handler {
	router := chi.NewRouter()

	router.Get("/", h.checkHealth)

	return router
}

// checkHealth ...
func (h *HealthHandler) checkHealth(w http.ResponseWriter, r *http.Request) {
	ServeJSON(w, "", http.StatusOK, "OK", nil, nil, nil)
}
