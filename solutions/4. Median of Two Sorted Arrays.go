func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}

	for i < len(nums1) {
		merged = append(merged, nums1[i])
		i++
	}
	for j < len(nums2) {
		merged = append(merged, nums2[j])
		j++
	}

	mid := len(merged) / 2
	if len(merged)%2 == 0 {
		return (float64(merged[mid-1]) + float64(merged[mid])) / 2.0
	} else {
		return float64(merged[mid])
	}
}