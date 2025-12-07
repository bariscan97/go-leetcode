type color int

func isIsomorphic(s string, t string) bool {
    colors := make(map[rune]color)
    sSlice := make([]color, len(s))
    var cnt color

    for idx, i := range s {
        if curr, ok := colors[i]; ok {
            sSlice[idx] = curr
        }else {
            cnt++
            colors[i] = cnt
            sSlice[idx] = colors[i] 
        }
    }
    
    
    clear(colors)
    cnt = 0
    
    for idx, i := range t {
        if _, ok := colors[i]; !ok {
           cnt++
           colors[i] = cnt
        }
        if sSlice[idx] != colors[i] {
            return false
        }
    }

    return true
}