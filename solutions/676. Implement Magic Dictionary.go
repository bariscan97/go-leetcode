type MagicDictionary struct {
	dic map[string]int
}

func Constructor() MagicDictionary {
	return MagicDictionary{
		dic: make(map[string]int),
	}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		var (
			size  int    = len(word)
			rWord []rune = []rune(word)
		)
		this.dic[word] = 1
		for i := 0; i < size; i++ {
			var prev rune = rWord[i]
			rWord[i] = '*'
			this.dic[string(rWord)]++
			rWord[i] = prev
		}
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	var (
		size  int    = len(searchWord)
		rWord []rune = []rune(searchWord)
		res   bool
	)

	for i := 0; i < size; i++ {
		var prev rune = rWord[i]
		rWord[i] = '*'
		if val, k := this.dic[string(rWord)]; k {
			if _, ok := this.dic[searchWord]; !ok || val != 1 {
				res = true
			}
		}
		rWord[i] = prev
	}

	return res
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */