func parse(num int) []int {
    var arr []int 
    for num > 0 {
        arr = append(arr, num % 10)
        num /= 10
    }
    return arr
}

func isPali(nums []int) bool {
    r := len(nums) -1
    l := 0
    for r > l {
        if nums[r] != nums[l] {
            return false
        }
        r--
        l++
    }
    return true
}

func isPalindrome(x int) bool {
    if 0 > x { return false }
    return isPali(parse(x))
}