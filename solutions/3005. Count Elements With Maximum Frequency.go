type Res struct {
	freq, cnt int
}

func maxFrequencyElements(nums []int) int {
	res := Res{}
	dic := make(map[int]int)
	for _, num := range nums {
		dic[num]++
		if dic[num] > res.freq {
			res = Res{
				freq: dic[num],
				cnt:  dic[num],
			}
		} else if dic[num] == res.freq {
			res.cnt += res.freq
		}
	}
	return res.cnt
}