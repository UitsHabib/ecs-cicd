package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/simple-web-app/service"
)

// Hanlder defines handler
type UserHandler struct {
	svc service.Service
}

// NewHandler ...
func NewUserHandler(svc service.Service) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

// Router ...
func (h *UserHandler) Router() http.Handler {
	router := chi.NewRouter()

	router.Get("/", h.getUsers)

	return router
}

// checkHealth ...
func (h *UserHandler) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.svc.GetUsers()
	if err != nil {
		ServeJSON(w, "", http.StatusInternalServerError, "Internal Server Error", nil, nil, err)
	}

	ServeJSON(w, "", http.StatusOK, "Successfully fetched", users, nil, nil)
}
