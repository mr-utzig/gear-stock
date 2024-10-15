package db

import (
	"database/sql"
	_ "embed"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func Connection() *sql.DB {
	var db *sql.DB
	var err error
	database := os.Getenv("TURSO_DATABASE_URL")
	token := os.Getenv("TURSO_AUTH_TOKEN")

	db, err = sql.Open("libsql", fmt.Sprintf("%s?authToken=%s", database, token))

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db conn: %s", err)
		os.Exit(1)
	}

	return db
}
