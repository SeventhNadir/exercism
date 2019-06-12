package letter

//ConcurrentFrequency counts word occurances in a given array of strings.
func ConcurrentFrequency(text []string) FreqMap {
	m := make(chan FreqMap)
	for _, s := range text {
		go func(s string) {
			m <- Frequency(s)
		}(s)
	}

	total := FreqMap{}
	for range text {
		for k, v := range <-m {
			total[k] += v
		}
	}
	return total
}
