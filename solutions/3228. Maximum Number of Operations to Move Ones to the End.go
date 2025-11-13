func maxOperations(s string) int {
    var (
        res, onesCnt int
        prev byte
    )

    for i := len(s) - 1; i > -1; i-- {
        if s[i] == '0' {
            if prev != '0' {
                onesCnt++ 
            }       
        }else {
            res += onesCnt
        }
        prev = s[i]
    }

    return res
}