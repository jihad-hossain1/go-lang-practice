package rest

import (
	"ecom/rest/handlers"
	"ecom/rest/handlers/user"
	"ecom/rest/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// root api
	mux.Handle("GET /", manager.With(http.HandlerFunc(handlers.Test)))

	// products api route
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts))) // route
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct)))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProduct)))
	mux.Handle("PUT /products/{id}", manager.With(http.HandlerFunc(handlers.UpdateProduct)))
	mux.Handle("DELETE /products/{id}", manager.With(http.HandlerFunc(handlers.DeleteProduct)))

	// users api route
	mux.Handle("POST /users", manager.With(http.HandlerFunc(user.CreateUser)))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(user.LoginUser)))

}
