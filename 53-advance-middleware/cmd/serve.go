package cmd

import (
	"ecom/middleware"
	"fmt"
	"net/http"
)

func Serve() {

	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux() // router
	initRoutes(mux, manager)

	fmt.Println("Server running on :3000")

	wrappedMux := manager.WrapMux(mux)

	err := http.ListenAndServe(":3000", wrappedMux) // nil

	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}
