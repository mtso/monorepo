package models

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const (
	CreateTableLeagues = `CREATE TABLE IF NOT EXISTS Leagues (
		    id varchar(24) NOT NULL UNIQUE PRIMARY KEY,
		    display_name varchar(256) NOT NULL,
		    created_at timestamp NOT NULL DEFAULT CURRENT_DATE
		)`

	InsertLeague = `INSERT INTO Leagues (id, display_name)
		VALUES ($1, $2)
		RETURNING id, display_name, created_at`
)

type League struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
}

func ConnectLeagues(db *sql.DB) (err error) {
	_, err = db.Exec(CreateTableLeagues)
	return
}

func CreateLeague(id, title interface{}) (l League, err error) {
	row := Conn.db.QueryRow(InsertLeague, id, title)

	err = row.Scan(&l.Id, &l.Title, &l.CreatedAt)
	return
}

func RandomId() string {
	now := time.Now()
	secret := fmt.Sprintf("%d%d", now.Unix(), now.UnixNano()) + os.Getenv("SESSION_SECRET")
	return HashId([]byte(secret))
}

func HashId(in []byte) string {
	h := sha256.Sum224(in)
	return fmt.Sprintf("%x", h)[:24]
}
