package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func NewTursoConn() *sql.DB {
	tursoUrl := os.Getenv("TURSO_DB_URL")
	tursoToken := os.Getenv("TURSO_DB_TOKEN")
	url := fmt.Sprintf("%s?authToken=%s", tursoUrl, tursoToken)

	db, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatal("Failed to open connection\n", err)
	}

	return db
}
