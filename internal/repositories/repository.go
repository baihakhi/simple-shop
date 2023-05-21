package repositories

import (
	"database/sql"
)

type repository struct {
	db *sql.DB
}

func InitRepository(db *sql.DB) Repositories {
	return &repository{
		db: db,
	}
}

type Repositories interface{}
