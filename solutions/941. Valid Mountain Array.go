func validMountainArray(arr []int) bool {
    color := 0
    
    for i := 0; i < len(arr)-1; i++ {
        if arr[i] < arr[i+1] {
            if color == 2 {
                return false
            }
            color = 1
        } else if arr[i] > arr[i+1] {
            if color == 0 {
                return false
            }
            color = 2
        } else {
            return false
        }
    }

    return color == 2
}
