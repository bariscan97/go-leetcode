type Res struct {
    Word string
    Cnt  int
}

func add(strWord string, dic map[string]int, res **Res) {
    if len(strWord) > 0 {
        dic[strWord] += 1
    }
    if *res == nil || dic[strWord] > (*res).Cnt {
        *res = &Res{
            Word: strWord,
            Cnt:  dic[strWord],
        }
    }
}

func mostCommonWord(paragraph string, banned []string) string {
    var (
        currWord []rune
        ord rune
        res *Res
    )
    
    dic := make(map[string]int)
    bans := make(map[string]bool)
    
    for _, i := range banned {
        bans[i] = true
    }
    
    for _, i := range paragraph + " " {
        ord = i
        if ord < 97 {
            ord += 32
        }
        if 97 <= ord && ord <= 122 {
            currWord = append(currWord, ord)
        }else {
            strWord := string(currWord)
            if !bans[strWord] {
                add(strWord, dic, &res)
            }
            currWord = []rune{}
        }
    }
    
    return res.Word
}