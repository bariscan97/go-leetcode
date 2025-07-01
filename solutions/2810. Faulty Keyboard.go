func Reverse(r []rune) []rune {
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return r
}

func finalString(s string) string {     
    res, r := []rune{}, []rune(s)
    for _, i := range r {
        if i == 'i' {
            Reverse(res)
        }else {
            res = append(res, i)
        }
    }
    return string(res)
}