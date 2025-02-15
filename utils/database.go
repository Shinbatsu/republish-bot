package utils

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetDB() *sql.DB {
	db, _ := sql.Open("sqlite3", "./mydatabase.db")
	return db
}

type Database struct {
	ptr *sql.DB
}

func (database *Database) GetDB() *sql.DB {
	database.ptr, _ = sql.Open("sqlite3", "./database.db")
	return database.ptr
}

func (database *Database) init() {

	database.ptr.Exec(`
		CREATE TABLE IF NOT EXISTS binds (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			source_id INTEGER NOT NULL,
			target_id INTEGER NOT NULL,
			name TEXT NOT NULL,
			bot_id INTEGER NOT NULL
		);
	`)
	database.ptr.Exec(`
		CREATE TABLE IF NOT EXISTS bots (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			user_name TEXT NOT NULL
		);
	`)
	database.ptr.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL
		);
	`)
	database.ptr.Exec(`
		CREATE TABLE IF NOT EXISTS chats (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL
		);
	`)
	database.ptr.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			target_id INTEGER NOT NULL,
			send_time DATETIME NOT NULL,
			post_id INTEGER NOT NULL
		);
	`)
	database.ptr.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			content JSON NOT NULL
		);
	`)
}
