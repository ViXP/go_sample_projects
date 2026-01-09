package data

import "database/sql"

type Store struct {
	DB     *sql.DB
	Models *Models
}
