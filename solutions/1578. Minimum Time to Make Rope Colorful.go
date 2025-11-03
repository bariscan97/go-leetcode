func minCost(colors string, neededTime []int) int {
    var (
        total int = neededTime[0]
        maxTime int = neededTime[0]
        prev byte = colors[0]
        res int
    )

    for i := 1; i < len(colors); i++ {
        if colors[i] == prev {
            maxTime = max(maxTime, neededTime[i])
            total += neededTime[i]
        }else {
            res += (total - maxTime)
            maxTime, total = neededTime[i], neededTime[i]
        }
        prev = colors[i]
    }
    
    return res + (total - maxTime)
}