package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Conn struct {
	db *sql.DB
}

// var dbconn *db.SQL

func Connect(connstring string) (*sql.DB, error) {
	conn, err := sql.Open("postgres", connstring)
	if err != nil {
		return nil, err
	}
	Conn.db = conn

	connects := []func(*sql.DB) error{
		ConnectLeagues,
	}

	for _, connect := range connects {
		err := connect(conn)
		if err != nil {
			return nil, err
		}
	}

	return conn, nil
}
