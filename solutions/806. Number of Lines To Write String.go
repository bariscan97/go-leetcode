func numberOfLines(widths []int, s string) []int {
    var (
        total int
        line  int = 1
    )
    
    for _, i := range s {
        val := widths[int(i - 97)]
        if val + total > 100 {
            total = 0
            line++ 
        }
        total += val
    }

    return []int{line, total}
}