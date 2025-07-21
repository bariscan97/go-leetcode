func makeFancyString(s string) string {
    var (
        prev rune = ' '
        cnt  int
        res  []rune
    )
    for _, curr := range s {
        if curr != prev {
            prev = curr
            cnt = 1
        }else {
            cnt++
        }
        if cnt < 3 {
            res = append(res, curr)
        }
    }
    return string(res) 
}