func halvesAreAlike(s string) bool {
    vowels := make(map[rune]bool)
    for _, i := range []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'} {
        vowels[i] = true
    }
    var (
        vowelsCnt int
        N         int = len(s)
        half      int = N / 2
    )
    for i := 0; i < half; i++ {
        if vowels[rune(s[i])] {
            vowelsCnt++
        }
        if vowels[rune(s[i + half])] {
            vowelsCnt--
        }
    }
	return vowelsCnt == 0
}