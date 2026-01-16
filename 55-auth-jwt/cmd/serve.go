package cmd

import (
	"ecom/config"
	"ecom/rest"
)

func Serve() {
	cnf := config.GetConfig()
	rest.Start(cnf)
}
