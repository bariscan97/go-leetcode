func Max(args ...int) int64 {
    var res int = -int(math.Pow(10,5))
    for _, arg := range args {
        if arg > res {
            res = arg
        }
    }
    return int64(res)
}

func maxProduct(nums []int) int64 {
    sort.Ints(nums)
    N := len(nums)
    return Max(
            nums[N - 3] * nums[N - 2] * int(math.Pow(10, 5)),
            nums[0] * nums[1] * int(math.Pow(10, 5)),
            nums[0] * nums[N - 1] * (int(math.Pow(10, 5)) * (-1)),
            nums[N - 1] * nums[N - 2] * int(math.Pow(10, 5)),
        )
}