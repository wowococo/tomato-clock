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
	db, err := sql.Open("sqlite3", "./TomatoClock.db")
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

	stmt, err := db.Prepare(task)
	hdlerr(err)
	stmt.Exec()

	tomato := `CREATE TABLE IF NOT EXISTS tomato(
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		taskID INTEGER, 
		duration INTEGER,
		progress REAL,
		startTime INTEGER,
		endTime INTEGER DEFAULT NULL, 
		updateTime INTEGER, 
		status TINYINT,
		FOREIGN KEY taskID REFERENCES task(id));`

	stmt, err = db.Prepare(tomato)
	hdlerr(err)
	stmt.Exec()

}

func insert(db *sql.DB, args ...interface{}) {
	for _, arg := range args {
		
	}

	 createTime := time.Now().Unix()
	stmt, err := db.Prepare(
		`INSERT INTO task(name, listID, status, createTime, updateTime) 
		values(?,?,?,?,?)`)
	hdlerr(err)
	stmt.Exec(stmt, name, 0, )
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
