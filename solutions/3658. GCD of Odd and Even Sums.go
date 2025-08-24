func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdOfOddEvenSums(n int) int {
    var odd, even int
    for i := 1; i < n * 2; i += 2 {
        odd += i
        even += i + 1
    }   
    return gcd(odd, even)
}