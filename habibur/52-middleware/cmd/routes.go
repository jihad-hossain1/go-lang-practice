package cmd

import (
	"ecom/handlers"
	"ecom/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /", manager.With(http.HandlerFunc(handlers.Test)))
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts))) // route
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct)))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProductById)))
}
