package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		 	log.Fatal("Error loading .env file")
	}
    fmt.Println("meu projeto em go esta rodando normalmente.")
}