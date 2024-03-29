package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const (
	CreateTablePlayers = `CREATE TABLE IF NOT EXISTS Players (
		id bigserial NOT NULL UNIQUE,
		league_id varchar(24) NOT NULL,
		display_name varchar(256) NOT NULL,
		elo int NOT NULL DEFAULT 1000
		)`

	SelectOrInsertPlayer = `WITH ExistingPlayer as (
			SELECT id, league_id, display_name, elo
			FROM Players
			WHERE league_id = $1
					AND display_name = $2
		), NewPlayer as (
			INSERT into Players (league_id, display_name)
			SELECT $1, $2
			WHERE NOT EXISTS (SELECT 1 from ExistingPlayer)
			returning id, league_id, display_name, elo
		)
		SELECT id, league_id, display_name, elo
		FROM NewPlayer
		union all select id, league_id, display_name, elo
		FROM ExistingPlayer`

	UpdatePlayerQuery = `UPDATE Players
		SET elo = $2
		WHERE id = $1`

	SelectPlayers = `SELECT display_name, elo
	    FROM Players
	    WHERE league_id = $1
	    ORDER BY elo DESC`
)

type Player struct {
	Id       int64  `json:"-"`
	LeagueId string `json:"-"`
	Username string `json:"username"`
	Elo      int64  `json:"elo"`
}

func ConnectPlayers(db *sql.DB) (err error) {
	_, err = db.Exec(CreateTablePlayers)
	return
}

func GetOrInsertPlayer(leagueid, name interface{}, tx ...*sql.Tx) (pl Player, err error) {
	var row *sql.Row
	if len(tx) > 0 {
		row = tx[0].QueryRow(SelectOrInsertPlayer, leagueid, name)
	} else {
		row = Conn.db.QueryRow(SelectOrInsertPlayer, leagueid, name)
	}
	err = row.Scan(&pl.Id, &pl.LeagueId, &pl.Username, &pl.Elo)
	return
}

func (pl *Player) Save(tx ...*sql.Tx) (err error) {
	if len(tx) > 0 {
		_, err = tx[0].Exec(UpdatePlayerQuery, pl.Id, pl.Elo)
	} else {
		_, err = Conn.db.Exec(UpdatePlayerQuery, pl.Id, pl.Elo)
	}
	return
}

func GetPlayers(id interface{}) ([]Player, error) {
	rows, err := Conn.db.Query(SelectPlayers, id)
	if err != nil {
		return nil, err
	}

	players := make([]Player, 0)
	var pl Player
	for rows.Next() {
		err := rows.Scan(&pl.Username, &pl.Elo)
		if err != nil {
			return nil, err
		}

		players = append(players, pl)
	}

	return players, nil
}
