func maximumCount(nums []int) int {
	neg := bisectLeft(nums, 0)
	pos := len(nums) - bisectRight(nums, 0)

	return max(neg, pos)
}

func bisectLeft(array []int, search int) int {
	low := 0
	high := len(array)

	for low < high {
		mid := low + (high-low)/2

		if array[mid] < search {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

func bisectRight(array []int, search int) int {
	low := 0
	high := len(array)

	for low < high {
		mid := low + (high-low)/2

		if array[mid] <= search {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}