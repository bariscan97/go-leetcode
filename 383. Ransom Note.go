func canConstruct(ransomNote string, magazine string) bool {
	counts := [26]int{}
	for i := 0; i < len(magazine); i++ {
		counts[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		counts[ransomNote[i]-'a']--
		if counts[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}