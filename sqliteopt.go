package main

import (
	"time"
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
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(255),
		listID INTEGER DEFAULT 0,
		status TINYINT,
		createTime INTEGER,
		updateTime INTEGER,
		finishTime INTEGER DEFAULT NULL);`

	_, err := db.Exec(task)
	hdlerr(err)

	tomato := `CREATE TABLE IF NOT EXISTS tomato(
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		taskID INTEGER, 
		duration INTEGER,
		startTime INTEGER,
		endTime INTEGER DEFAULT NULL, 
		updateTime INTEGER, 
		status TINYINT,
		FOREIGN KEY taskID REFERENCES task(id));`

	_, err = db.Exec(tomato)
	hdlerr(err)	

}

func insert() {
	 createTime := time.Now().Unix()
	 stmt, err := "INSERT INTO task(name, listID, status, createTime) values()"

}

func query() {

}

func hdlerr(err error) {
	if err != nil {
		panic(err)
	}
}
