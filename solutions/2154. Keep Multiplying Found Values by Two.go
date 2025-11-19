func findFinalValue(nums []int, original int) int {
    dic := make(map[int]bool)
    for _, i := range nums {
        dic[i] = true
    }
    var f func(num int) int
    f = func(num int) int {
        if !dic[num] {
            return num
        } 
        return f(num * 2)
    }
    return f(original)
}