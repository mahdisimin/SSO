package main

import (
	"SSO/config"
	"SSO/delivery/router"
)

func main() {
	err := config.LoadDotEnv(".env")
	if err != nil {
		panic(err)
	}

	router.RouteResolver()
}
