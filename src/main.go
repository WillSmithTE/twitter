package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/willsmithte/twitter/src/musicCharts"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	musicCharts.Main()
}
