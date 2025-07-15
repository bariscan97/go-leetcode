func isValid(word string) bool {
    if len(word) < 3 { return false }
    vowels := map[rune]bool {
        'a': true,
        'e': true,
        'i': true,
        'o': true,
        'u': true,
    }
    var a, b bool
    for _, i := range word {
        if 'A' <= i && i <= 'Z' || 'a' <= i && i <= 'z' {
            if vowels[i] || vowels[i-32] || vowels[i + 32] {
                a = true
            }else {
                b = true
            }
        }else if !('0' <= i && i <= '9') {
            return false
        }
    }
    return a && b 
}