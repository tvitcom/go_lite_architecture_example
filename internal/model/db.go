package model

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type DB struct {
    *sql.DB
}
type MongoDB struct {
    *sql.DB
}

func NewDB(drv, dataSourceName string) (*DB, error) {
    db, err := sql.Open(drv, dataSourceName)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return &DB{db}, nil
}
func NewMongoDB(drv, dataSourceName string) (*MongoDB, error) {
    db, err := sql.Open(drv, dataSourceName)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return &MongoDB{mongodb}, nil
}
