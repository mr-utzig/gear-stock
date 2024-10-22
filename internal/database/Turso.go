package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var Turso *sql.DB

func NewTursoConn() {
	tursoUrl := os.Getenv("TURSO_DATABASE_URL")
	tursoToken := os.Getenv("TURSO_AUTH_TOKEN")
	url := fmt.Sprintf("%s?authToken=%s", tursoUrl, tursoToken)

	db, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatal("Failed to open connection\n", err)
	}

	Turso = db
}
