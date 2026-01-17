package cmd

import (
	"ecom/config"
	"ecom/rest"
	"ecom/rest/handlers/product"
	"ecom/rest/handlers/user"
	"ecom/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler()

	server := rest.NewServer(cnf, productHandler, userHandler)

	server.Start()
}
