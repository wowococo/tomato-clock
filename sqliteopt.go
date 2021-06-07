package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbName = "TomatoClock"
)

func init() {
	db := createdb()
	createtb(db)
}

func createdb() *sql.DB {
	db, err := sql.Open("sqlite3", "TomatoClock.db")
	hdlerr(err)
	return db
}

func createtb(db *sql.DB) {
	task := `CREATE TABLE IF NOT EXISTS task(
		id integer primary key autoincrement,
		name varchar(255),
		listID integer,
		status tinyint,
		createTime integer,
		updateTime integer,
		finishTime integer);`

	_, err := db.Exec(task)
	hdlerr(err)

	tomato := `CREATE TABLE IF NOT EXISTS tomato(
		id integer primary key autoincrement, 
		taskID integer, 
		duration integer,
		startTime integer,
		endTime integer, 
		updateTime integer, 
		status tinyint
		FOREIGN KEY taskID REFERENCES task(id));`

	_, err = db.Exec(tomato)
	hdlerr(err)	

}

// func insert() {

// }

// func query() {

// }

func hdlerr(err error) {
	if err != nil {
		panic(err)
	}
}
