func numWaterBottles(numBottles int, numExchange int) int {
    var res, a, b int = numBottles, 0, 0
    for {
        a = numBottles % numExchange
        b = numBottles / numExchange
        res += b
        numBottles = a + b
        if numBottles < numExchange {
            break
        }
    }
    return res
}