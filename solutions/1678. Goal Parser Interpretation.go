func interpret(command string) string {
    var (
        stack []int
        res   []rune
    )
    for idx, i := range command {
        switch i {
            case '(':
                stack = append(stack, idx)
            case ')':
                last := stack[len(stack) - 1]
                stack = stack[:len(stack) - 1]
                if idx - last == 1 {
                    res = append(res, 'o')
                }
            default:
                res = append(res, i)
        }
    }
    return string(res)
}