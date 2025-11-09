func arrangeCoins(n int) int {
    var num int
    for n > num {
        num++
        n -= num
    }
    return num
}