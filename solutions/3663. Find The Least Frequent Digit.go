type Res struct {
    Num int
    Cnt int
}

func getLeastFrequentDigit(n int) int {
    dic := make(map[int]int)
    res := Res{ Num: 11, Cnt: 99999999}
	
    for n > 0 {
		digit := n % 10
        dic[digit]++
        n /= 10		
	}

    for k, v := range dic {
        if v < res.Cnt || (k < res.Num && res.Cnt == v)  {
            res = Res{
                Num: k,
                Cnt: v,
            }
        }
    }
    
    return res.Num
}