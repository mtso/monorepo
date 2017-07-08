package models

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

const (
	CreateTableGames = `CREATE TABLE IF NOT EXISTS Games (
		id bigserial NOT NULL UNIQUE,
		league_id varchar(24) NOT NULL,
		winner_id bigint NOT NULL,
		loser_id bigint NOT NULL,
		created_at timestamp NOT NULL DEFAULT NOW()
		)`

	InsertGame = `IF EXISTS (SELECT 1 FROM Leagues WHERE id = $1) THEN
			INSERT INTO Games (league_id, winner_id, loser_id)
			VALUES ($1, $2, $3)
			RETURNING id, created_at
		END IF`

	// SelectGame = `SELECT DISTINCT ON (Games.id)
	// 	Games.id,
	// 	Games.created_at,
	// 	Players.id,
	// 	Players.

	// 	FROM Games, Players
	// 	`
)

type Game struct {
	Id        int64     `json:"-"`
	CreatedAt time.Time `json:"createdAt"`

	League League `json:"-"`
	Winner Player `json:"winner,omitempty"`
	Loser  Player `json:"loser,omitempty"`
}

func ConnectGames(db *sql.DB) (err error) {
	_, err = db.Exec(CreateTableGames)
	return
}

func AddGame(leagueid, winnername, losername interface{}, calcHandler func(Player, Player) int64) (gm Game, err error) {
	_, err = GetLeagueById(leagueid)
	if err != nil {
		return
	}

	tx, err := Conn.db.Begin()
	if err != nil {
		return gm, err
	}

	// if league exists:
	// get winner and loser players by leagueid
	//
	// calculate transfer points
	//

	winner, err := GetOrInsertPlayer(leagueid, winnername, tx)
	if err != nil {
		return
	}
	loser, err := GetOrInsertPlayer(leagueid, losername, tx)
	if err != nil {
		return
	}

	pointsTransferred := calcHandler(winner, loser)
	winner.Elo += pointsTransferred
	loser.Elo -= pointsTransferred

	err = winner.Save(tx)
	if err != nil {
		return
	}
	err = loser.Save(tx)
	if err != nil {
		return
	}

	row := tx.QueryRow(InsertGame, leagueid, winner.Id, loser.Id)
	err = row.Scan(&gm.Id, &gm.CreatedAt)
	if err != nil {
		return
	}

	gm.Winner = winner
	gm.Loser = loser
	return
}
