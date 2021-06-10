package sqliteopt

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

const (
	dbName = "TomatoClock"
)

func init() {
	db = createdb()
	createtb()
}

func createdb() *sql.DB {
	db, err := sql.Open("sqlite3", "./TomatoClock.db")
	hdlerr(err)
	return db
}

func createtb() {
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
		timefocused INTEGER,
		progress REAL,
		startTime INTEGER,
		endTime INTEGER DEFAULT NULL, 
		updateTime INTEGER, 
		status TINYINT,
		FOREIGN KEY(taskID) REFERENCES task(id));`

	stmt, err = db.Prepare(tomato)
	hdlerr(err)
	stmt.Exec()

}

func insert(statement string, args ...interface{}) int64 {
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	hdlerr(err)

	res, err := stmt.Exec(args...)
	hdlerr(err)

	id, err := res.LastInsertId()
	hdlerr(err)

	return id

}

func InsertTask(args ...interface{}) int64 {
	statement := `INSERT INTO task(name, listID, status, createTime, updateTime) 
			values(?, ?, ?, ?, ?)`
	return insert(statement, args...)
}

func InsertTomato(taskID int64, args ...interface{}) int64 {
	s := fmt.Sprintf("values(%v, ?, ?, ?, ?, ?, ?)", taskID)
	statement := `INSERT INTO tomato(
			taskID, duration, timefocused, progress, 
			startTime, updateTime, status) ` + s

	return insert(statement, args...)

}

func UpdateTomato(args ...interface{}) int64 {
	statement := `UPDATE tomato SET timefocused=? 
			AND progress=? AND endTime=? AND updateTime=? 
			AND status=?
			WHERE id=?`
	stmt, err := db.Prepare(statement)
	hdlerr(err)

	res, err := stmt.Exec(args...)
	hdlerr(err)

	affect, err := res.RowsAffected()
	hdlerr(err)

	return affect
}

func Query() {
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
