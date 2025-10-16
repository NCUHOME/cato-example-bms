package main

import (
	"log"

	"cato-example-bms/internal/config/app"
	_ "cato-example-bms/internal/handlers/books"
)

func main() {
	srv := app.GetApp()
	if err := srv.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
