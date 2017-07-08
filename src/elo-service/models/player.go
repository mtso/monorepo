package models

const (
	CreateTablePlayers = `CREATE TABLE IF NOT EXISTS Players (
		id bigserial NOT NULL UNIQUE,
		display_name NOT NULL,
		league_id varchar(24) NOT NULL,
		elo int NOT NULL
		)`
)
