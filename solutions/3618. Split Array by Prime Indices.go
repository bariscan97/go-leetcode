func splitArray(nums []int) int64 {
    n := len(nums)
    
    isPrime := make([]bool, n)
    if n > 2 {
        for i := 2; i < n; i++ {
            isPrime[i] = true
        }
        limit := int(math.Sqrt(float64(n - 1)))
        for p := 2; p <= limit; p++ {
            if isPrime[p] {
                for multiple := p * p; multiple < n; multiple += p {
                    isPrime[multiple] = false
                }
            }
        }
    }

    var sumA, sumB int64
    for i, v := range nums {
        if i < len(isPrime) && isPrime[i] {
            sumA += int64(v)
        } else {
            sumB += int64(v)
        }
    }

    diff := sumA - sumB
    if diff < 0 {
        diff = -diff
    }
    return diff
}