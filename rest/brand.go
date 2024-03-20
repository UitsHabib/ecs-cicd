package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/simple-web-app/service"
)

// Hanlder defines handler
type BrandHandler struct {
	svc service.Service
}

// NewHandler ...
func NewBrandHandler(svc service.Service) *BrandHandler {
	return &BrandHandler{
		svc: svc,
	}
}

// Router ...
func (h *BrandHandler) Router() http.Handler {
	router := chi.NewRouter()

	router.Get("/", h.getBrands)

	return router
}

// checkHealth ...
func (h *BrandHandler) getBrands(w http.ResponseWriter, r *http.Request) {
	brands, err := h.svc.GetBrands()
	if err != nil {
		ServeJSON(w, "", http.StatusInternalServerError, "Internal Server Error", nil, nil, err)
	}

	ServeJSON(w, "", http.StatusOK, "Successfully fetched", brands, nil, nil)
}
