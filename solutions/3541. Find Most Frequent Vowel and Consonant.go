type Dic map[rune]int

func maxFreqSum(s string) int {
    dic := make(Dic)
    var vowelCnt, notVowelCnt int

    for _, i := range s {
        dic[i]++
        switch i {
        case 'a', 'e', 'i', 'o', 'u':
            vowelCnt = max(vowelCnt, dic[i])
        default:
            notVowelCnt = max(notVowelCnt, dic[i])
        }
    }

    return vowelCnt + notVowelCnt
}