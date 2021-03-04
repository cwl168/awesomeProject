package main

import (
	"fmt"
	"runtime"
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
func ConcurrentFrequency(s []string) FreqMap {
	var (
		result  = FreqMap{}
		channel = make(chan FreqMap, len(s))
	)

	for _, v := range s {
		go func(text string) {
			channel <- Frequency(text)
		}(v)
	}
	//遍历map，取值，非遍历通道吧
	for range s {
		for k, v := range <-channel {
			result[k] += v
		}
	}

	/*for f := range channel {
		for k, v := range f {
			result[k] += v
		}
	}*/

	return result
}
func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
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
	con := ConcurrentFrequency([]string{euro, dutch, us})
	fmt.Println(con)
}
