package main

import (
	routes "clean-golang/application/router"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("server start on port 9000")
	routes.RunServer()
}
