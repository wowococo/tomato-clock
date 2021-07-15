package sqliteopt

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

const (
	dbName = "TomatoClock.db"
)

func init() {
	db = createdb()
	createtable()
}

func createdb() *sql.DB {
	gopath := os.Getenv("GOPATH")
	dir := gopath + "/src/tomato-clock"
	dataSrc := fmt.Sprintf("%s/%s", dir, dbName)

	db, err := sql.Open("sqlite3", dataSrc)
	hdlerr(err)
	return db
}

func createtable() {
	task := `CREATE TABLE IF NOT EXISTS task(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(255) UNIQUE,
		listID INTEGER DEFAULT 0,
		status TINYINT,
		createTime INTEGER,
		updateTime INTEGER,
		endTime INTEGER DEFAULT NULL);`

	stmt, err := db.Prepare(task)
	hdlerr(err)
	defer stmt.Close()

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
	defer stmt.Close()

	stmt.Exec()

}

func insert(statement string, args ...interface{}) int64 {
	stmt, err := db.Prepare(statement)
	hdlerr(err)
	defer stmt.Close()

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

func queryTask(name string) (id int64) {
	statement := `SELECT id FROM task WHERE name = ?`
	rows, err := db.Query(statement, name)
	hdlerr(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id)
		hdlerr(err)
	}
	return

}

func updateTask(args ...interface{}) (affect int64) {
	var statement string
	if len(args) == 2 {
		statement = `UPDATE task SET updateTime=? WHERE id=?`
	}
	if len(args) == 4 {
		statement = `UPDATE task SET endTime=?, updateTime=?, status=? WHERE id=?`
	}

	stmt, err := db.Prepare(statement)
	hdlerr(err)
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	hdlerr(err)

	affect, err = res.RowsAffected()
	hdlerr(err)

	return
}

func updateTomato(args ...interface{}) int64 {
	statement := `UPDATE tomato SET timefocused=?, progress=?, endTime=?, updateTime=?, status=? 
		WHERE id=?`
	stmt, err := db.Prepare(statement)
	hdlerr(err)
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	hdlerr(err)

	affect, err := res.RowsAffected()
	hdlerr(err)

	return affect
}

const (
	lc  = "lineChart"
	txt = "text"
)

type Metric string

// metrics needed to query
func (mtc Metric) Query(table, col, timeslot string) string {
	var statement string
	switch table {
	case "tomato":
		if col == "progress" {
			// coalesce accepts at least two arguments and return the first non-null value, to avoid sum(progress) is NULL
			statement = "SELECT COALESCE(SUM(progress), 0) FROM tomato WHERE status in (1,3)"
		}
		if col == "timefocused" {
			statement = "SELECT COALESCE(SUM(timefocused), 0) FROM tomato WHERE 1=1"
		}
	case "task":
		statement = "SELECT COALESCE(COUNT(id), 0) FROM task WHERE status = 1"
	}
	res, ok := _query(txt, statement, timeslot).(float64)
	if !ok {
		// optimize
		return ""
	}
	prec := 1
	switch table {
	case "task":
		prec = 0
	case "tomato":
		if col == "timefocused" {
			res = res / (60 * 60)
		}
	}
	return strconv.FormatFloat(res, 'f', prec, 64)
}

type LChart string

// linechart inputs needed to query
func (lct LChart) Query(table, timeslot string) interface{} {
	var statement string
	switch table {
	case "tomato":
		switch timeslot {
		case "untiltoday":
			statement = "SELECT DATE(endTime, 'unixepoch') day, SUM(progress) count FROM tomato WHERE status in (1, 3)"
		case "untilweek":
			statement = "SELECT strftime('%Y-%W', endTime, 'unixepoch') week, SUM(progress) count FROM tomato WHERE status in (1,3)"
		case "untilmonth":
			statement = "SELECT strftime('%Y-%m', endTime, 'unixepoch') month, SUM(progress) count FROM tomato WHERE status in (1,3)"
		}
	case "task":
		switch timeslot {
		case "untiltoday":
			statement = "SELECT DATE(endTime, 'unixepoch') day, COUNT(id) count FROM task WHERE status = 1"
		case "untilweek":
			statement = "SELECT strftime('%Y-%W', endTime, 'unixepoch') week, COUNT(id) count FROM task WHERE status = 1"
		case "untilmonth":
			statement = "SELECT strftime('%Y-%m', endTime, 'unixepoch') month, COUNT(id) count FROM task WHERE status = 1"
		}
	}

	res := _query(lc, statement, timeslot)
	return res
}

func _query(chartType, statement, timeslot string) interface{} {
	now := time.Now()
	y, M, d, location := now.Year(), now.Month(), now.Day(), now.Location()
	switch timeslot {
	default:
		statement += ";"
	case "thisweek":
		weekday := now.Weekday()
		st := now.AddDate(0, 0, int(time.Monday-weekday))
		// Sunday as the last day
		et := now.AddDate(0, 0, int(time.Sunday+7-weekday))
		st = time.Date(st.Year(), st.Month(), st.Day(), 0, 0, 0, 0, st.Location())
		et = time.Date(et.Year(), et.Month(), et.Day(), 24, 0, 0, 0, et.Location())

		statement += fmt.Sprintf(" and endTime >= %v and endTime < %v;", st.Unix(), et.Unix())
	case "today":
		// y, M, d, location := now.Year(), now.Month(), now.Day(), now.Location()
		st := time.Date(y, M, d, 0, 0, 0, 0, location)
		et := time.Date(y, M, d+1, 0, 0, 0, 0, location)
		statement += fmt.Sprintf(" and endTime >= %v and endTime < %v;", st.Unix(), et.Unix())
	case "untiltoday":
		// y, M, d, location := now.Year(), now.Month(), now.Day(), now.Location()
		st := time.Date(y, M-1, d, 0, 0, 0, 0, location)
		et := now
		statement += fmt.Sprintf(" and endTime >= %v and endTime <= %v GROUP BY day;", st.Unix(), et.Unix())
	case "untilweek":
		mon := func(t time.Time) time.Time {
			weekday := t.Weekday()
			mondate := t.AddDate(0, 0, int(time.Monday-weekday))
			return mondate
		}
		mondate := mon(now)
		y, M, d, location := mondate.Year(), mondate.Month(), mondate.Day(), mondate.Location()
		st := mon(time.Date(y, M-6, d, 0, 0, 0, 0, location))
		et := now
		statement += fmt.Sprintf(" and endTime >= %v and endTime <= %v GROUP BY week;", st.Unix(), et.Unix())
	case "untilmonth":
		// y, M, location := now.Year(), now.Month(), now.Location()
		st := time.Date(y-1, M, 1, 0, 0, 0, 0, location)
		et := now
		statement += fmt.Sprintf(" and endTime >= %v and endTime <= %v GROUP BY month;", st.Unix(), et.Unix())
	}

	rows, err := db.Query(statement)
	hdlerr(err)
	defer rows.Close()
	var (
		res     float64
		dt      string
		count   float64
		dtCount = make(map[string]float64)
	)
	switch chartType {
	case txt:
		for rows.Next() {
			err = rows.Scan(&res)
			hdlerr(err)
		}
		return res
	case lc:
		for rows.Next() {
			err = rows.Scan(&dt, &count)
			hdlerr(err)
			dtCount[dt] = count
		}
		return dtCount
	default:
		return nil
	}
}

func hdlerr(err error) {
	if err != nil {
		panic(err)
	}
}
