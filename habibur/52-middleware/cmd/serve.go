package cmd

import (
	"ecom/global_router"
	"ecom/handlers"
	"ecom/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux() // router

	mux.Handle("GET /", middleware.Logger(http.HandlerFunc(handlers.Test)))
	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts)) // route
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductById))

	fmt.Println("Server running on :3000")

	globalRouter := global_router.GlobalRouter(mux)
	err := http.ListenAndServe(":3000", globalRouter) // nil

	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}
