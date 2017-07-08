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

	InsertGame = `INSERT INTO Games (league_id, winner_id, loser_id)
			VALUES ($1, $2, $3)
			RETURNING id, created_at`

	SelectGames = `SELECT
		games.id,
		games.created_at,
		winner.display_name,
		winner.elo,
		loser.display_name,
		loser.elo
		FROM Games
		JOIN Players as winner on Games.winner_id = winner.id 
		JOIN Players as loser on Games.loser_id = loser.id 
		WHERE Games.league_id = $1
		ORDER BY Games.created_at DESC`
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

	err = tx.Commit()
	if err != nil {
		return
	}

	gm.Winner = winner
	gm.Loser = loser
	return
}

func GetGames(id interface{}) ([]Game, error) {
	rows, err := Conn.db.Query(SelectGames, id)
	if err != nil {
		return nil, err
	}

	games := make([]Game, 0)
	var g Game
	for rows.Next() {
		err := rows.Scan(&g.Id, &g.CreatedAt,
			&g.Winner.Username, &g.Winner.Elo,
			&g.Loser.Username, &g.Loser.Elo)
		if err != nil {
			return nil, err
		}

		games = append(games, g)
	}

	return games, nil
}
