package repository

import "database/sql"

type CRUD interface {
	Create(tx *sql.Tx) (int, error)
	// Read()
	// Update(*sql.Tx)
	// Delete(*sql.Tx)
}

func CreateInterface(s CRUD, tx *sql.Tx) (int, error) {
	return s.Create(tx)
}
