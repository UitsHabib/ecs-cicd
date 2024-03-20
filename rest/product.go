package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/simple-web-app/service"
)

// Hanlder defines handler
type ProductHandler struct {
	svc service.Service
}

// NewHandler ...
func NewProductHandler(svc service.Service) *ProductHandler {
	return &ProductHandler{
		svc: svc,
	}
}

// Router ...
func (h *ProductHandler) Router() http.Handler {
	router := chi.NewRouter()

	router.Get("/", h.getProducts)

	return router
}

// checkHealth ...
func (h *ProductHandler) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.svc.GetProducts()
	if err != nil {
		ServeJSON(w, "", http.StatusInternalServerError, "Internal Server Error", nil, nil, err)
	}

	ServeJSON(w, "", http.StatusOK, "Successfully fetched", products, nil, nil)
}
