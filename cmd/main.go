package main

import (
	"fmt"
	"os"

	"github.com/Valgard/godotenv"
	"github.com/mr-utzig/gear-stock/app"
)

func main() {
	dotenv := godotenv.New()
	if err := dotenv.Load(".env"); err != nil {
		fmt.Fprintf(os.Stderr, "failed to load .env: %s", err)
		os.Exit(1)
	}

	app.Setup()
}
