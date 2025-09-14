func earliestTime(tasks [][]int) int {
    var res int = 10e9
    for _, task := range tasks {
        res = min(res, task[0] + task[1])
    }
    return res
}