func generateParenthesis(n int) []string {
    var (
        dfs func(curr string, left, right int)
        res []string 
    )

    dfs = func(curr string, left, right int) {
        if left == n && right == n {
            res = append(res, curr)
        }
        if left < n {
            dfs(curr + "(", left + 1, right)
        }
        if right < left {
            dfs(curr + ")", left, right + 1)
        }
    }
    
    dfs("", 0, 0)

    return res
}