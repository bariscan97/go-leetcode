func possibleStringCount(word string) int {
    cnt := 1
    prev := word[0]
    res := 0

    for i := 1; i < len(word); i++ {
        if word[i] == prev {
            cnt++
        } else {
            res += cnt - 1
            cnt = 1
        }
        prev = word[i]
    }

    res += cnt

    return res
}