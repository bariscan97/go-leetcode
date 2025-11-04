func equalFrequency(word string) bool {
	dic := make(map[rune]int)
	freq := make(map[int]int)

	for _, i := range word {
		dic[i]++
	}

	for _, v := range dic {
		freq[v]++
	}

	arr := reflect.ValueOf(freq).MapKeys()
	a := arr[0].Interface().(int)

	if len(freq) == 1 {
		return freq[a] == 1 || freq[1] > 0
	}

	b := arr[1].Interface().(int)

	if len(freq) > 2 {
		return false
	}

	if freq[1] == 1 {
		return true
	}

	if (a > b && a-b > 1) || (b > a && b-a > 1) {
		return false
	}

	if (a > b && freq[a] != 1) || (b > a && freq[b] != 1) {
		return false
	}

	return true
}