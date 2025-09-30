type Res struct {
    groupSize int   
    str       []rune
}

func majorityFrequencyGroup(s string) string {
    
    dic := make(map[rune]int)
    groups := make(map[int][]rune)
    res := Res{}
    
    for _, i := range s {
        dic[i]++
    }
    
    for k, v := range dic {
        groups[v] = append(groups[v], k)
        if len(groups[v]) > len(res.str) ||
            (len(groups[v]) == len(res.str) && v > res.groupSize) {
                res.str = groups[v]
                res.groupSize = v
        }
    }
    
    return string(res.str)
}