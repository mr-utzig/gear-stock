package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
	"github.com/mr-utzig/gear-stock/internal/api"
	"github.com/mr-utzig/gear-stock/internal/database"
)

func main() {
	port := flag.String("port", ":6969", "The port which the webserver will listen.")
	flag.Parse()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	database.NewTursoConn()
	defer database.Turso.Close()

	server := api.NewServer(*port)

	server.Start()
}
