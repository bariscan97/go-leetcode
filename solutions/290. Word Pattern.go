type Color int

func colorPattern(s string) []Color {
    var color Color
    size := len(s)
    pat := make([]Color, size)
    dic := make(map[rune]Color)
    for idx, i := range s {
        if _, ok := dic[i]; !ok {
            dic[i] = color
            color++
        }
        pat[idx] = dic[i]
    }
    return pat
}
 
func sPattern(s string) []Color {
    var color Color
    var curr []byte
    N := len(s)
    pat := []Color{}
    dic := make(map[string]Color)
    for i := 0; i < N + 1; i++ {
        if i == N || s[i] == ' ' {
            str_curr := string(curr)
            if _, ok := dic[str_curr]; !ok {
                dic[str_curr] = color
                color++
            }
            pat = append(pat, dic[str_curr])
            curr = []byte{}
        }else{
            curr = append(curr, s[i])
        }
    }
    return pat
}

func wordPattern(pattern string, s string) bool {
    a := colorPattern(pattern)
    b := sPattern(s)
    aSize := len(a)
    bSize := len(b)
    if aSize != bSize {
        return false
    }
    for i := 0; i < aSize; i++ {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}   