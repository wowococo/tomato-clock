package sqliteopt

import (
	"database/sql"
	"fmt"
	"time"

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

func insertTask(args ...interface{}) int64 {
	statement := `INSERT INTO task(name, listID, status, createTime, updateTime) 
			values(?, ?, ?, ?, ?)`
	return insert(statement, args...)
}

func insertTomato(taskID int64, args ...interface{}) int64 {
	s := fmt.Sprintf("values(%v, ?, ?, ?, ?, ?, ?)", taskID)
	statement := `INSERT INTO tomato(
			taskID, duration, timefocused, progress, 
			startTime, updateTime, status) ` + s

	return insert(statement, args...)

}

func updateTomato(args ...interface{}) int64 {
	statement := `UPDATE tomato SET timefocused=?, progress=?, endTime=?, updateTime=?, status=? 
		WHERE id=?`
	stmt, err := db.Prepare(statement)
	hdlerr(err)

	res, err := stmt.Exec(args...)
	hdlerr(err)

	affect, err := res.RowsAffected()
	hdlerr(err)

	return affect
}

func QueryTomato(total, thisweek, today bool) {
	statement := "SELECT SUM(progress) FROM tomato WHERE status in（1,3)"
	if total {
		statement += ";"
	}
	if thisweek {
		weekday := time.Now().Weekday()
		st := time.Now().AddDate()
		et := time.Now().AddDate()
		statement += fmt.Sprintf(" and endTime between %v and %v;", st, et)
	}
	if today {
		statement += " and endTime between x and y;"
	}
	Query(statement)
}

func Query(statement string) {
	rows, err := db.Query(statement)
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
