package models

import ()

const (
	CreateTableGames = `CREATE TABLE IF NOT EXISTS Games (
		id bigserial NOT NULL UNIQUE,
		league_id varchar(24) NOT NULL,
		winner_id bigint NOT NULL,
		loser_id bigint NOT NULL
		)`
)
