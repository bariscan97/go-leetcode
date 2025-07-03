func kthCharacter(k int) byte {
    var chars []rune = []rune{'a'}
    for {
        size := len(chars)
        if size > k { break }
        for i := 0; i < size; i++ {
            chars = append(chars, ((chars[i] + 1) % 97) + 97)
        }
    }
    return byte(chars[k - 1])
}