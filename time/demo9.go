package main

import (
	"fmt"
	"time"
)

func main() {
	var duration_Seconds time.Duration = (1250 * 10) * time.Millisecond
	var duration_Minute time.Duration = 2 * time.Minute

	dur, _ := time.ParseDuration("1m30s")
	fmt.Printf("dur [%v],[%d],[%T]\n", dur, dur, dur)
	var float64_Seconds float64 = duration_Seconds.Seconds()
	var float64_Minutes float64 = duration_Minute.Minutes()

	fmt.Printf("Seconds [%.3f]\nMinutes [%.2f]\n", float64_Seconds, float64_Minutes)
}
