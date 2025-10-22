package main

import (
	"log"
	"tusur-forms/internal/app"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
