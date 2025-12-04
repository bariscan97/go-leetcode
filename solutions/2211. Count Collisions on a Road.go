func countCollisions(directions string) int {
    var R, S, res int

    for _, dir := range directions {
        switch dir {
            case 'L':
                if R > 0 {
                    res += 1 + R
                    R, S = 0, 1
                }else if S > 0 {
                    res += 1
                    R, S = 0, 1
                }
            case 'R':
                S, R = 0, R + 1
            default:
                res += R
                R, S = 0, 1
        }
    }
    
    return res
}