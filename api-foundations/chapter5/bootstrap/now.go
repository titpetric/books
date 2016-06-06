package bootstrap;

import "time"

var StartTime float64;
func Now() float64 {
	myTime := float64(time.Now().UnixNano()) / 1000000.0;
	if (StartTime < 0.000001) {
		StartTime = myTime;
	}
	return myTime - StartTime;
}
