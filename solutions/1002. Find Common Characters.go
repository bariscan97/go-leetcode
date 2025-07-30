func Counter(word string) map[rune]int {
    dic := make(map[rune]int)
    for _, i := range word {
        dic[i]++
    }
    return dic
}

func min(args ...int) int {
    val := args[0]
    for _, i := range args[1:] {
        if i < val {
            val = i
        }
    }
    return val
}

type Val struct {
    Cnt int
    Ok  int
}

func commonChars(words []string) []string {
    size := len(words)
    common := make(map[rune]Val)
    for _, word := range words {
        for i, j := range Counter(word) {
            if v, ok := common[i]; ok {
                v.Cnt = min(v.Cnt, j)
                v.Ok++
                common[i] = v 
            }else {
                common[i] = Val{
                    Cnt: j,
                    Ok: 1,
                }
            }
        }
    }
    var res []string
    for k, v := range common {
        if v.Ok == size {
            for i := 0; i < v.Cnt; i++ {
                res = append(res, string(k))
            }
        }
    }
    
    return res
}