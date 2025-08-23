func majorityElement(nums []int) []int {
    var res []int
    N := len(nums)
    dic := make(map[int]int)
    for _, num := range nums {
        dic[num]++
        if dic[num] == N / 3 + 1 {
            res = append(res, num)
        }
    }
    return res
}