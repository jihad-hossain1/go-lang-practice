package cmd

import (
	"ecom/global_router"
	"ecom/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(middleware.MdlwrOne, middleware.Logger)

	mux := http.NewServeMux() // router
	initRoutes(mux, manager)
	fmt.Println("Server running on :3000")

	globalRouter := global_router.GlobalRouter(mux)
	err := http.ListenAndServe(":3000", globalRouter) // nil

	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}
