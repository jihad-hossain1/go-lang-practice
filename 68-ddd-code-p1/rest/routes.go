package rest

import (
	"ecom/rest/handlers"
	"ecom/rest/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// root api
	mux.Handle("GET /", manager.With(http.HandlerFunc(handlers.Test)))

}
