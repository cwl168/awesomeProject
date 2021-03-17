package timeDuration

import (
	"fmt"
	"time"
)

func Duration1() {
	var waitFiveHundredMillisections int64 = 500 //500毫秒

	startingTime := time.Now().UTC()
	time.Sleep(10 * time.Millisecond) //10毫秒
	endingTime := time.Now().UTC()

	var duration time.Duration = endingTime.Sub(startingTime) //Duration 类型中时间的基本单位是 Nanosecond
	var durationAsInt64 = int64(duration)                     //将一个表示 10 毫秒的 Duration 类型对象转换为 int64 类型时，我实际上得到的是 10,000,000

	fmt.Println(durationAsInt64)

	if durationAsInt64 >= waitFiveHundredMillisections {
		fmt.Printf("Time Elapsed : Wait[%d] Duration[%d]\n", waitFiveHundredMillisections, durationAsInt64)
	} else {
		fmt.Printf("Time DID NOT Elapsed : Wait[%d] Duration[%d]\n", waitFiveHundredMillisections, durationAsInt64)
	}

	u, _ := time.ParseDuration("1s")
	fmt.Printf("One second is %d microseconds.\n", u.Microseconds()) //One second is 1000000 microseconds.
}
func Duration2() {
	var waitFiveHundredMillisections time.Duration = 500 * time.Millisecond //500000000 纳秒

	startingTime := time.Now().UTC()
	time.Sleep(10 * time.Millisecond)
	endingTime := time.Now().UTC()

	var duration time.Duration = endingTime.Sub(startingTime)

	if duration >= waitFiveHundredMillisections {
		fmt.Printf("Time Elapsed : Wait[%d] Duration[%d]\n", waitFiveHundredMillisections, duration)
	} else {
		fmt.Printf("Time DID NOT Elapsed : Wait[%d] Duration[%d]\n", waitFiveHundredMillisections, duration)
	}

	var duration_Milliseconds time.Duration = 500 * time.Millisecond
	var duration_Seconds time.Duration = (1250 * 10) * time.Millisecond
	var duration_Minute time.Duration = 2 * time.Minute

	fmt.Printf("Milli [%v]\nSeconds [%v]\nMinute [%v]\n", duration_Milliseconds, duration_Seconds, duration_Minute)

	//不可以
	/*connectTimeout := 10
	time.Sleep(time.Second * connectTimeout)*/

	//需要类型转化
	//connectTimeout := 10
	//time.Sleep(time.Duration(connectTimeout) * time.Second)

	// Golang 和时间相关的可以直接使用数字， 但是不能使用float 浮点类型， 也不能直接是数值型变量
	//yourTime := 2
	//time.Sleep(1 * time.Second)
	//time.Sleep(1.1 * time.Second)
	//time.Sleep(time.Duration(yourTime) * time.Second)
	//time.Sleep(yourTime * time.Second)

	d, _ := time.ParseDuration("1m30s") //900 0000 0000 纳秒
	fmt.Printf("%v,%d，%T\n", d, d, d)

	str := "2018-01-14T21:45:54+08:00"
	//先将时间转换为字符串

	tt, _ := time.Parse("2006-01-02T15:04:05Z07:00", str)

	//格式化时间
	fmt.Println(tt.Format("2006-01-02 15:04:05"))

}
