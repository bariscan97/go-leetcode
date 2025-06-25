func capitalizeTitle(title string) string {
	arr := strings.Split(title, " ")
	res := ""
	for index, word := range arr {
		if len(word) > 2 {
			res += strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		} else {
			res += strings.ToLower(word)
		}
		if len(arr)-1 > index {
			res += " "
		}
	}
	return res
}