package service

import "time"

var StartTime float64

// return time in milliseconds (ms, us=ms/1000, ns=us/1000)
func Now() float64 {
	myTime := float64(time.Now().UnixNano()) / 1000000.0
	if StartTime < 0.000001 {
		StartTime = myTime
	}
	return myTime - StartTime
}
