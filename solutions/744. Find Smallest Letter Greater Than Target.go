func nextGreatestLetter(letters []byte, target byte) byte {
    left, right := 0, len(letters)-1
    res := letters[0] 

    for left <= right {
        mid := left + (right-left)/2

        if letters[mid] > target {
            res = letters[mid] 
            right = mid - 1    
        } else {
            left = mid + 1
        }
    }

    return res
}
