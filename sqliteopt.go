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
	db, err := sql.Open("sqlite3", "./TomatoClock.db")
	hdlerr(err)
	return db
}

func createtb(db *sql.DB) {
	task := `CREATE TABLE IF NOT EXISTS task(
		id integer primary key autoincrement,
		name varchar(255),
		listID integer,
		status tinyint,
		createTime text,
		updateTime text,
		finishTime text);`

	stmt, err := db.Prepare(task)
	hdlerr(err)
	stmt.Exec()

	tomato := `CREATE TABLE IF NOT EXISTS tomato(
		id integer primary key autoincrement, 
		taskID integer, 
		duration integer,
		progress real,
		startTime text,
		endTime text, 
		updateTime text, 
		status tinyint,
		FOREIGN KEY(taskID) REFERENCES task(id));`

	stmt, err = db.Prepare(tomato)
	hdlerr(err)
	stmt.Exec()

}

func insert(db *sql.DB) {
	stmt, err := db.Prepare("insert into task values()")
	hdlerr(err)
	stmt.Exec()
}

func update(db *sql.DB) {

}

func query(db *sql.DB) {

	rows, err := db.Query("select sum(progress) from tomato where status in (1, 2)")
	defer rows.Close()
	hdlerr(err)
	for rows.Next() {
		err = rows.Scan()
		hdlerr(err)
	}
}

func hdlerr(err error) {
	if err != nil {
		panic(err)
	}
}
