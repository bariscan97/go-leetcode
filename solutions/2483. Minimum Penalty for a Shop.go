type P struct {
    y, n int
}

type Res struct {
    cnt, idx int
}

func bestClosingTime(customers string) int {
    var (
        nCount, yCount int
        N int = len(customers)
    )
    
    prefix := make([]*P, N)
    
    for i := 0; i < N; i++ {
        if customers[i] == 'N' {
            nCount++ 
        }else {
            yCount++
        }
        prefix[i] = &P{
            n: nCount,
            y: yCount,
        }
    }

    res:= Res {
        cnt: nCount,
        idx: N,
    }
    
    for i := 0; i < N; i++ {
        n := prefix[i].n
        y := prefix[N - 1].y - prefix[i].y
        if customers[i] == 'Y' {
            y++
        }else {
            n--
        }
        if n + y < res.cnt || (res.cnt == n + y && i < res.idx) {
            res = Res{
                cnt: n + y,
                idx: i,
            }
        }
    }
    return res.idx
}