package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/xvbnm48/go-medium-ecomerce/route"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	fmt.Println("main application starts")
	loadEnv()

	log.Fatal(route.RunAPI(":8080"))
}
