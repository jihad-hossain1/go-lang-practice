package rest

import (
	"ecom/config"
	"ecom/rest/handlers/product"
	"ecom/rest/handlers/user"
	"ecom/rest/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	cnf            *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(
	cnf *config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler,
) *Server {
	return &Server{
		cnf,
		productHandler,
		userHandler,
	}
}

func (server *Server) Start() {

	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux() // router
	wrappedMux := manager.WrapMux(mux)

	// initRoutes(mux, manager)
	server.productHandler.ProductRoutes(mux, manager)
	server.userHandler.UserRoutes(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("server running on port", addr)

	err := http.ListenAndServe(addr, wrappedMux) // nil

	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}

}
