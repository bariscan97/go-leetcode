func sumOfEncryptedInt(nums []int) int {
	var res int
	for _, num := range nums {
		var max_char float64
		str_num := strconv.Itoa(num)
		for _, i := range str_num {
			max_char = math.Max(max_char, float64(i))
		}
		val, _ := strconv.Atoi(string(rune(max_char)))
		start := 1
		for digit := 0; digit < len(str_num); digit++ {
			res += start * val
			start *= 10
		}
	}
	return res
}