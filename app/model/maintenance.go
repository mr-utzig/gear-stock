package model

import "database/sql"

type Maintenance struct {
}

func (m *Maintenance) Create(db *sql.DB) error {
	return nil
}
