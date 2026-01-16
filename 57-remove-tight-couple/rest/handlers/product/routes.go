package product

import (
	"ecom/rest/middleware"
	"net/http"
)

func (h *Handler) ProductRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	// products api route
	mux.Handle("GET /products", manager.With(http.HandlerFunc(h.GetProducts))) // route
	mux.Handle("POST /products", manager.With(
		http.HandlerFunc(h.CreateProduct),
		middleware.AuthJwt,
	))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(h.GetProduct)))
	mux.Handle("PUT /products/{id}", manager.With(http.HandlerFunc(h.UpdateProduct)))
	mux.Handle("DELETE /products/{id}", manager.With(http.HandlerFunc(h.DeleteProduct)))

}
