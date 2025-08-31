func recoverOrder(order []int, friends []int) []int {
    dic := make(map[int]bool)
    res := []int{}
    
    for _, friend := range friends {
        dic[friend] = true
    }

    for _, i := range order {
        if dic[i] {
            res = append(res, i)
        }
    }

    return res
}