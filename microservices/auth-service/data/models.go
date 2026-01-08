package data

import (
	"database/sql"
	"time"
)

const timeout = time.Second * 5

var dbPool *sql.DB

type Models struct {
	User *User
}

func NewModels(pool *sql.DB) *Models {
	dbPool = pool

	return &Models{
		User: &User{},
	}
}
