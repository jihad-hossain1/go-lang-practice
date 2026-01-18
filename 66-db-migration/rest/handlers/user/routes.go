package user

import (
	"ecom/rest/middleware"
	"net/http"
)

func (h *Handler) UserRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// users api route
	mux.Handle("POST /users", manager.With(http.HandlerFunc(h.CreateUser)))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(h.LoginUser)))

}
