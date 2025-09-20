func letterCasePermutation(s string) []string {
    var (
        res []string
        f   func(i int, curr []byte)
    )

    f = func(i int, curr []byte) {
        if i == len(s) {
            res = append(res, string(curr))
            return
        }
        if '0' <= s[i] && s[i] <= '9' {
            f(i + 1, append(curr, s[i]))
        }else {
            if 'a' <= s[i] && s[i] <= 'z' {
                f(i + 1, append(curr, s[i] - 32))
            }else {
                f(i + 1, append(curr, s[i] + 32))
            }
            f(i + 1, append(curr, s[i]))
        }
    }
    
    f(0, []byte{})
    
    return res
}