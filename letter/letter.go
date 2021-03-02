package letter

// FreqMap records the frequency of each rune in a given text.
// FreqMap is a map that holds the rune as the key and number of occurences as the value.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
// Frequency takes a string and figures out the frequency of those words.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency takes a slice of strings
// and concurrently finds the frequency of those words.
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

	for range s {
		for k, v := range <-channel {
			result[k] += v
		}
	}

	return result
}

//map[10:21 32:119 33:1 39:5 44:16 45:1 46:3 59:2 63:2 65:2 66:1 68:2 69:2 70:2 71:2 72:3 75:1 77:2 78:1 79:5 80:1 84:1(T) 87:6 90:1 97:47(a) 98:16 99:7 100:32 101:100 102:10 103:22 104:34 105:42 106:5 107:7 108:32 109:10 110:50 111:29 112:7 114:50 115:38 116:55(t) 117:19 118:11 119:13 121:9 235:1 246:2 252:2]