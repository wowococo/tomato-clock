package sqliteopt

import (
	"time"
)

func PostTask(name string, status int8) int64 {
	createTime, updateTime := getTime()
	// listID is reserved field
	listID := 0

	id := insertTask(name, listID, status, createTime, updateTime)
	return id
}

func GetTask(name string) (id int64, ok bool) {
	id = queryTask(name)
	ok = true
	if id == int64(0) {
		ok = false
	}
	return
}

func PutTask(args ...interface{}) int64 {
	_, ut := getTime()
	//add an element at the beginning of the slice
	args = append([]interface{}{ut}, args...)
	affect := updateTask(args...)

	return affect
}

func PostTomato(taskID int64, d time.Duration, status int8) int64 {
	startTime, updateTime := getTime()
	// the beginning of the tomato, focus, progress are 0
	focus, progress := 0, 0

	id := insertTomato(taskID, d.Seconds(), focus, progress, startTime, updateTime, status)
	return id
}

func PutTomato(id int64, timeleft, d time.Duration, status int8) int64 {
	// update tomato(timefocused, progress, endTime, updateTime, status)
	endTime, updateTime := getTime()
	focus := (d - timeleft)
	progress := float64(focus) / float64(d)
	// if status == 0 || status == 2 {
	// fix: endTime should be updated
	// }
	affect := updateTomato(focus.Seconds(), progress, endTime, updateTime, status, id)
	return affect
}

func getTime() (st, ut int64) {
	st = time.Now().Unix()
	ut = st
	return
}
