func checkString(s string) bool {
    var flag bool
    for _, i := range s {
        if i == 'b' {
            flag = true
        }
        if flag && i == 'a' {
            return false
        }
    }
    return true
}