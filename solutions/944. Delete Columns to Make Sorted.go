func minDeletionSize(strs []string) int {
    var (
        res int
        flag bool
    )
    
    for i := 0; i < len(strs[0]); i++ {
        flag = false
        for j := 1; j < len(strs); j++ {
            if strs[j - 1][i] > strs[j][i] {
                flag = true
                break
            }
        }
        if flag { res++ }
    }

    return res
}