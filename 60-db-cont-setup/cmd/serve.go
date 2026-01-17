package cmd

import (
	"ecom/config"
	"ecom/repo"
	"ecom/rest"
	"ecom/rest/handlers/product"
	"ecom/rest/handlers/user"
	"ecom/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()

	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(cnf, productHandler, userHandler)

	server.Start()
}
