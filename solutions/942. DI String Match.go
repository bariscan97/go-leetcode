func diStringMatch(s string) []int {
    n := len(s)
    res := make([]int, n+1) 
    low, high := 0, n

    for i, char := range s {
        if char == 'I' {
            res[i] = low
            low++
        } else {
            res[i] = high
            high--
        }
    }
    
    res[n] = low
    
    return res
}