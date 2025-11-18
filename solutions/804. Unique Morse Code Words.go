func uniqueMorseRepresentations(words []string) int {
    morseAlphabet := []string{
        ".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--..",
    }
    
    var curr string
    seens := make(map[string]bool)
    
    for _, word := range words {
        curr = ""
        for _, i := range word {
                curr += morseAlphabet[int(i - 97)]
        }
        seens[curr] = true
    }
    
    return len(seens)
}