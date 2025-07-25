func intersection(nums1 []int, nums2 []int) []int {
    dic := make(map[int]bool)
    var res []int
    for _, num := range nums1 {
        dic[num] = true
    }
    for _, num := range nums2 {
        if dic[num] {
            res = append(res, num)
            dic[num] = false
        }
    }
    return res
}