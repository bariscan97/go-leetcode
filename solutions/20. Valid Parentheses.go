func isValid(s string) bool {
    var stack []rune
    dic := map[rune]rune{
            ')': '(',
            '}': '{',
            ']': '[',
        }
    for _, i := range s {
        if val, ok := dic[i]; ok && len(stack) != 0 && stack[len(stack) - 1] == val {
            stack = stack[:len(stack) - 1]
        }else {
            stack = append(stack, i)
        }
    }
    return len(stack) == 0
}