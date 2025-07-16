func Separate(n int) []int {
    var arr []int
    for n > 0 {
        arr = append([]int{n % 10}, arr...)
        n /= 10
    }
    return arr
}

func separateDigits(nums []int) []int {
    var res []int
    for _, num := range nums {
        res = append(res, Separate(num)...)
    }
    return res
}