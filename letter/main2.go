package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// sema limits the number of goroutines for frequency counting to 10.
// 我们可以用一个有容量限制的buffered channel来控制并发，这类似于操作系统里的计数信号量概念。
var sema = make(chan struct{}, 10)

//限制并发
func ConcurrentFrequency4(s []string) FreqMap {
	var (
		result  = FreqMap{}
		channel = make(chan FreqMap, len(s))
	)

	for _, v := range s {
		sema <- struct{}{}
		go func(text string) {
			defer func() { <-sema }()
			channel <- Frequency(text)
		}(v)
	}

	for range s {
		for k, v := range <-channel {
			result[k] += v
		}
	}

	return result
}

//没有限制goroutine
func ConcurrentFrequency3(s []string) FreqMap {
	var (
		result  = FreqMap{}
		channel = make(chan FreqMap, len(s))
	)

	for _, v := range s {
		go func(text string) {
			channel <- Frequency(text)
		}(v)
	}

	for range s {
		for k, v := range <-channel {
			result[k] += v
		}
	}

	return result
}

func ConcurrentFrequency2(s []string) FreqMap {
	var wg sync.WaitGroup
	wg.Add(len(s))
	var (
		result  = FreqMap{}
		channel = make(chan FreqMap, len(s))
	)

	for _, v := range s {
		sema <- struct{}{}
		go func(text string) {
			defer wg.Done()
			defer func() { <-sema }()
			channel <- Frequency(text)
		}(v)
	}
	wg.Wait()
	close(channel)
	for f := range channel {
		for k, v := range f {
			result[k] += v
		}
	}

	return result
}

func ConcurrentFrequency1(s []string) FreqMap {
	var wg sync.WaitGroup
	wg.Add(len(s))
	var (
		result  = FreqMap{}
		channel = make(chan FreqMap, len(s))
	)

	for _, v := range s {
		sema <- struct{}{} // acquire a token
		go func(text string) {
			defer wg.Done()
			defer func() { <-sema }() // release the token
			channel <- Frequency(text)
		}(v)
	}
	go func() {
		wg.Wait()
		close(channel)
	}()
	//channel 为buffer channel,所以会出现 channel关闭后，for range 还在循环从channe读取数据，但是channel数据长度肯定是 小于等于 len(s)

	//channel close后，发送数据将会触发 panic 异常，还能接受channel中的数据，等channel中的数据为空，返回nil, 接受就会停止了
	//channel 没有close的话，如果channel没数据的话，接受只会阻塞，等待着发送者发送数据,如果数据满了，发送数据端将阻塞等待接收
	for f := range channel {
		for k, v := range f {
			result[k] += v
		}
	}
	return result
}

func ConcurrentFrequency6(s []string) FreqMap {
	var wg sync.WaitGroup
	wg.Add(len(s))
	var (
		result  = FreqMap{}
		channel = make(chan FreqMap)
	)

	for _, v := range s {
		go func(text string) {
			defer wg.Done()
			channel <- Frequency(text)
		}(v)
	}
	go func() {
		wg.Wait()
		close(channel)
	}()

	for range s {
		for k, v := range <-channel {
			result[k] += v
		}
	}

	return result
}

//无缓冲通道
func ConcurrentFrequency5(s []string) FreqMap {
	var (
		result  = FreqMap{}
		channel = make(chan FreqMap)
	)

	for _, v := range s {
		go func(text string) {
			channel <- Frequency(text)
		}(v)
	}

	for range s {
		for k, v := range <-channel {
			result[k] += v
		}
	}
	close(channel)

	return result
}
func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("\nthe number of goroutines: ", runtime.NumGoroutine())
	}()

	var (
		euro = `Freude schöner Götterfunken
Tochter aus Elysium,
Wir betreten feuertrunken,
Himmlische, dein Heiligtum!
Deine Zauber binden wieder
Was die Mode streng geteilt;
Alle Menschen werden Brüder,
Wo dein sanfter Flügel weilt.`

		dutch = `Wilhelmus van Nassouwe
ben ik, van Duitsen bloed,
den vaderland getrouwe
blijf ik tot in den dood.
Een Prinse van Oranje
ben ik, vrij, onverveerd,
den Koning van Hispanje
heb ik altijd geëerd.`

		us = `O say can you see by the dawn's early light,
What so proudly we hailed at the twilight's last gleaming,
Whose broad stripes and bright stars through the perilous fight,
O'er the ramparts we watched, were so gallantly streaming?
And the rockets' red glare, the bombs bursting in air,
Gave proof through the night that our flag was still there;
O say does that star-spangled banner yet wave,
O'er the land of the free and the home of the brave?`
	)
	//方法一
	start := time.Now()
	ret1 := ConcurrentFrequency1([]string{euro, dutch, us, dutch, us, euro, euro, euro, euro, dutch, us, dutch, euro, dutch, us, dutch, dutch, us, dutch, euro, dutch, us, dutch, us, dutch, dutch, us, dutch, euro, dutch, us, dutch, euro, dutch, us, dutch})
	cost := time.Since(start)
	fmt.Printf("ConcurrentFrequency1 cost=[%s]\n", cost)

	//方法二
	start2 := time.Now()
	_ = ConcurrentFrequency2([]string{euro, dutch, us, dutch, us, euro, euro, euro, euro, dutch, us, dutch, euro, dutch, us, dutch, dutch, us, dutch, euro, dutch, us, dutch, us, dutch, dutch, us, dutch, euro, dutch, us, dutch, euro, dutch, us, dutch})
	cost2 := time.Since(start2)
	fmt.Printf("ConcurrentFrequency2 cost=[%s]\n", cost2)

	//方法三
	start3 := time.Now()
	_ = ConcurrentFrequency3([]string{euro, dutch, us, dutch, us, euro, euro, euro, euro, dutch, us, dutch, euro, dutch, us, dutch, dutch, us, dutch, euro, dutch, us, dutch, us, dutch, dutch, us, dutch, euro, dutch, us, dutch, euro, dutch, us, dutch})
	cost3 := time.Since(start3)
	fmt.Printf("ConcurrentFrequency3 cost=[%s]\n", cost3)

	//方法四
	start4 := time.Now()
	ret4 := ConcurrentFrequency4([]string{euro, dutch, us, dutch, us, euro, euro, euro, euro, dutch, us, dutch, euro, dutch, us, dutch, dutch, us, dutch, euro, dutch, us, dutch, us, dutch, dutch, us, dutch, euro, dutch, us, dutch, euro, dutch, us, dutch})
	cost4 := time.Since(start4)
	fmt.Printf("ConcurrentFrequency4 cost=[%s]\n", cost4)

	//方法五
	start5 := time.Now()
	_ = ConcurrentFrequency5([]string{euro, dutch, us, dutch, us, euro, euro, euro, euro, dutch, us, dutch, euro, dutch, us, dutch, dutch, us, dutch, euro, dutch, us, dutch, us, dutch, dutch, us, dutch, euro, dutch, us, dutch, euro, dutch, us, dutch})
	cost5 := time.Since(start5)
	fmt.Printf("ConcurrentFrequency5 cost=[%s]\n", cost5)

	////方法6
	start6 := time.Now()
	_ = ConcurrentFrequency6([]string{euro, dutch, us, dutch, us, euro, euro, euro, euro, dutch, us, dutch, euro, dutch, us, dutch, dutch, us, dutch, euro, dutch, us, dutch, us, dutch, dutch, us, dutch, euro, dutch, us, dutch, euro, dutch, us, dutch})
	cost6 := time.Since(start6)
	fmt.Printf("ConcurrentFrequency6 cost=[%s]\n", cost6)

	for k, v := range ret4 {
		fmt.Printf("%v:%v\t", string(k), v)
	}
	fmt.Println()
	fmt.Println()

	for k, v := range ret1 {
		fmt.Printf("%v:%v\t", string(k), v)
	}
}

/*func FormatPrint(ret FreqMap )  {
	fmt.Println()
	var keys []int32
	for k := range ret {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for k, v := range ret {
		fmt.Printf("%v:%v\t", string(k), v)
	}
	fmt.Println()
}*/
