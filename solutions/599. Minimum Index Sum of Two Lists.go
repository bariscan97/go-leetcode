type Res struct {
	Arr    []string
	IdxSum int
}

func findRestaurant(list1 []string, list2 []string) []string {
	res := Res{
		Arr:    []string{},
		IdxSum: int(^uint(0) >> 1),
	}
	dic := make(map[string]int)
	for i, word := range append(list1, list2...) {
		if idx, ok := dic[word]; ok {
			total := idx + i
			if total < res.IdxSum {
				res = Res{
					Arr:    []string{word},
					IdxSum: total,
				}
			} else if total == res.IdxSum {
				res.Arr = append(res.Arr, word)
			}
		}
		dic[word] = i
	}
	return res.Arr
}