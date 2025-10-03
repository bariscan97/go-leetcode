func reverseOnlyLetters(s string) string {
    N := len(s)
    runes := []rune(s)
    l, r := 0, N -1
    for {
        for l < N && !unicode.IsLetter(runes[l]) {
            l++
        }
        for r > -1 && !unicode.IsLetter(runes[r]) {
            r--
        }
        if l >= r {
            break
        }
        runes[l], runes[r] = runes[r], runes[l]
        l++
        r--
        
    }
    return string(runes)
}