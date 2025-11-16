func abs(num int) int {
    if num >= 0 {
        return num
    }
    return -num
}

func minLengthAfterRemovals(s string) int {
    var a, b int 
    for _, i := range s {
        if i == 'a' {
            a++
        }else{
            b++
        }
    }
    return abs(a - b)
}