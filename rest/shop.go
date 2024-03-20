package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/simple-web-app/service"
)

// Hanlder defines handler
type ShopHandler struct {
	svc service.Service
}

// NewHandler ...
func NewShopHandler(svc service.Service) *ShopHandler {
	return &ShopHandler{
		svc: svc,
	}
}

// Router ...
func (h *ShopHandler) Router() http.Handler {
	router := chi.NewRouter()

	router.Get("/", h.getShops)

	return router
}

// checkHealth ...
func (h *ShopHandler) getShops(w http.ResponseWriter, r *http.Request) {
	shops, err := h.svc.GetShops()
	if err != nil {
		ServeJSON(w, "", http.StatusInternalServerError, "Internal Server Error", nil, nil, err)
	}

	ServeJSON(w, "", http.StatusOK, "Successfully fetched", shops, nil, nil)
}
