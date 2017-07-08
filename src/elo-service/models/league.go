package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	CreateTableLeagues = `CREATE TABLE IF NOT EXISTS Leagues (
		    id varchar(24) NOT NULL UNIQUE PRIMARY KEY,
		    display_name varchar(256) NOT NULL,
		    created_at timestamp NOT NULL DEFAULT CURRENT_DATE 
		)`
)

func ConnectLeague(db *db.Conn) {
	db.Exec(CreateTableLeague)
}

func HashId(in []byte) string {
	h := sha256.Sum224(in)
	return fmt.Sprintf("%x", h)[:24]
}
