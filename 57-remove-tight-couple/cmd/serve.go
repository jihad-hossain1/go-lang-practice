package cmd

import (
	"ecom/config"
	"ecom/rest"
	"ecom/rest/handlers/product"
	"ecom/rest/handlers/user"
)

func Serve() {
	cnf := config.GetConfig()

	productHandler := product.NewHandler()
	userHandler := user.NewHandler()

	server := rest.NewServer(cnf, productHandler, userHandler)

	server.Start()
}
