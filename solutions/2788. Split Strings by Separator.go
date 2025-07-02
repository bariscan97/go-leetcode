func splitWordsBySeparator(words []string, separator byte) []string {
	var res []string
	var s rune = rune(separator)
	for _, word := range words {
		var c []rune
		for _, i := range word {
			if i == s {
				if len(c) > 0 {
					res = append(res, string(c))
					c = []rune{}
				}
			} else {
				c = append(c, i)
			}
		}
		if len(c) > 0 {
			res = append(res, string(c))
		}
	}
	return res
}