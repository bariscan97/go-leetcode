func minimumMoves(s string) int {
    var (
        res, cnt int
        flag bool
    )

    for _, i := range s {
        if i == 'X' {
            flag = true
        }
        if flag {
            cnt++
        }
        
        if cnt == 3 {
            cnt = 0
            res++
            flag = false
        }
    }
    
    if cnt > 0 { res++ }
    
    return res
}