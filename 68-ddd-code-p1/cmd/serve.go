package cmd

import (
	"ecom/config"
	"ecom/infra/db"
	"ecom/repo"
	"ecom/rest"
	prdHanler "ecom/rest/handlers/product"
	usrHandler "ecom/rest/handlers/user"
	middleware "ecom/rest/middleware"
	"ecom/user"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Migrate DB
	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	// domains
	usrSvc := user.NewService(userRepo)
	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := prdHanler.NewHandler(middlewares, productRepo)
	userHandler := usrHandler.NewHandler(cnf, usrSvc)

	server := rest.NewServer(cnf, productHandler, userHandler)

	server.Start()
}
