package main

import (
	"os"

	"github.com/mktbsh/golang-devcontainer/router"
)

func main() {
	rt := router.Init()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	rt.Logger.Fatal(rt.Start(":" + port))
}
