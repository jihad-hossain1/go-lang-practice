package rest

import (
	"ecom/config"
	"ecom/rest/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Start(cnf config.Config) {

	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux() // router
	initRoutes(mux, manager)

	wrappedMux := manager.WrapMux(mux)

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("server running on port", addr)

	err := http.ListenAndServe(addr, wrappedMux) // nil

	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}

}
