func intersect(nums1 []int, nums2 []int) []int {
	i, j := 0, 0
	n := len(nums1)
	m := len(nums2)
	sort.Ints(nums1)
	sort.Ints(nums2)
	var ans []int
	for i < n && j < m {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			ans = append(ans, nums1[i])
			i++
			j++
		}
	}
	return ans
}