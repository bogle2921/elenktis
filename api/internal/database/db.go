package database

import (
	"database/sql"

	"github.com/bogle2921/elenktis/api/internal/utils"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("pgx", utils.DB_URL)
	if err != nil {
		panic("Could not connect to database " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `create table if not exists users (
	id SERIAL PRIMARY KEY,
	admin BOOLEAN NOT NULL,
	username VARCHAR(32) NOT NULL UNIQUE,
	password VARCHAR(64) NOT NULL,
	salt bytea
	)`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table")
	}

	createSystemsTable := `create table if not exists systems (
		id SERIAL PRIMARY KEY,
		ip VARCHAR(15) NOT NULL,
		friendlyname VARCHAR(64),
		services TEXT []
		)`

	_, err = DB.Exec(createSystemsTable)
	if err != nil {
		panic("Could not create systems table")
	}
}
