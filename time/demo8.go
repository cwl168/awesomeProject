package main

import (
	"fmt"
	"time"
)

func main() {
	var duration_Milliseconds time.Duration = 500 * time.Millisecond
	var duration_Seconds time.Duration = (1250 * 10) * time.Millisecond
	var duration_Minute time.Duration = 2 * time.Minute

	fmt.Printf("Milli [%v]\nSeconds [%v]\nMinute [%v]\n", duration_Milliseconds, duration_Seconds, duration_Minute)

	//500 毫秒
	//var waitFiveHundredMillisections int64 = 500
	var waitFiveHundredMillisections time.Duration = 500 * time.Millisecond

	//时间  2020-12-04 06:44:18.22467 +0000 UTC
	startingTime := time.Now().UTC()
	//睡眠10毫秒
	time.Sleep(10 * time.Millisecond)
	endingTime := time.Now().UTC()

	var duration time.Duration = endingTime.Sub(startingTime)
	var durationAsInt64 = int64(duration)

	if durationAsInt64 >= waitFiveHundredMillisections {
		fmt.Printf("Time Elapsed : Wait[%d] Duration[%d]\n", waitFiveHundredMillisections, durationAsInt64)
	} else {
		fmt.Printf("Time DID NOT Elapsed : Wait[%d] Duration[%d]\n", waitFiveHundredMillisections, durationAsInt64)
	}

}
