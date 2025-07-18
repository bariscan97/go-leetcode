type Phone map[rune]string

var dic Phone = make(Phone)

func letterCombinations(digits string) []string {
	dic['2'] = "abc"
	dic['3'] = "def"
	dic['4'] = "ghi"
	dic['5'] = "jkl"
	dic['6'] = "mno"
	dic['7'] = "pqrs"
	dic['8'] = "tuv"
	dic['9'] = "wxyz"

	var (
		res []string
		f   func(curr string, i int)
	)

	f = func(curr string, i int) {
		if i == len(digits) {
			if curr != "" {
				res = append(res, curr)
			}
			return
		}
		for _, v := range dic[rune(digits[i])] {
			f(curr+string(v), i+1)
		}
	}

	f("", 0)

	return res
}