package sqliteopt

import "time"


func PostTask(name string, status int8) int64 {
	createTime, updateTime := getTime()
	// listID is reserved field
	listID:= 0

	id := insertTask(name, listID, status, createTime, updateTime)
	return id
}

func PostTomato(taskID int64, d time.Duration, status int8) int64 {
	startTime, updateTime := getTime()
	// the beginning of the tomato, focus, progress are 0
	focus, progress := 0, 0
	
	id := insertTomato(taskID, d, focus, progress, startTime, updateTime, status)
	return id
}

func PutTomato(id int64, timeleft, d time.Duration, status int8) int64 {
	// update tomato(timefocused, progress, endTime, updateTime, status)
	endTime, updateTime := getTime()
	focus := (d - timeleft)
	progress := focus / d

	affect := updateTomato(focus, progress, endTime, updateTime, status, id)
	return affect
}

func getTime() (st, ut int64) {
	st = time.Now().Unix()
	ut = st
	return 
}