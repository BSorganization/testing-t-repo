package main

import (
	"log"

	"github.com/golang_buffalo/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
