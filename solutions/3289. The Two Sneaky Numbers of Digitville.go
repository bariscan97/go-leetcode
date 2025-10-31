func getSneakyNumbers(nums []int) []int {
    res := []int{}
    dic := make(map[int]bool)
    
    for _, num := range nums {
        if dic[num] {
            res = append(res, num)
        }
        dic[num] = true
    }
    
    return res
}