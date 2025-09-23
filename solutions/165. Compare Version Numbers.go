func parse(s string) []int {
	var res []int
	for _, i := range strings.Split(s, ".") {
		num, _ := strconv.Atoi(i)
		res = append(res, num)
	}
	return res
}

func compareVersion(version1 string, version2 string) int {
	v1, v2 := parse(version1), parse(version2)
	v1Size, v2Size := len(v1), len(v2)
	var i, j int
	for i < v1Size || j < v2Size {
		val1, val2 := 0, 0
		if i < v1Size {
			val1 = v1[i]
		}
		if j < v2Size {
			val2 = v2[j]
		}
		if val1 > val2 {
			return 1
		}
		if val1 < val2 {
			return -1
		}
		i++
		j++
	}

	return 0
}