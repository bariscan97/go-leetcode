func minimumAbsDifference(arr []int) [][]int {
    dic := make(map[int][][]int)
    mini := 1_000_000_000_7
    
    slices.Sort(arr)

    for i := 1; i < len(arr); i++ {
        diff := arr[i] - arr[i - 1]
        dic[diff] = append(
            dic[diff],
            []int{arr[i - 1], arr[i]},
        )
        mini = min(mini, diff)
    }

    return dic[mini]
}