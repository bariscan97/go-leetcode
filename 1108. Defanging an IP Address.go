func defangIPaddr(address string) string {
	res := ""
	for _, i := range address {
		if string(i) == "." {
			res += "[.]"
		} else {
			res += string(i)
		}
	}
	return res
}