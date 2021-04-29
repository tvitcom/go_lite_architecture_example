package models

import (
	"database/sql"
)

type SetupStorage struct {
	DB *sql.DB
}
