func uncommonFromSentences(s1 string, s2 string) []string {
    
	dic := make(map[string]int)
    
	for _, i := range append(strings.Split(s1, " "), strings.Split(s2, " ")...) {
        dic[i]++
        
    }  
    
    var res []string
    
    for k, v := range dic {
        if v == 1 {
            res = append(res, k)
        }
    }

    return res
}