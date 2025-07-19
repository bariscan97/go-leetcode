func split(s string) []string {
	var (
		val []rune
		res []string
	)
	for _, i := range s {
		if i != '/' {
			val = append(val, i)
		} else {
			res = append(res, string(val))
			val = []rune{}
		}
	}
	return append(res, string(val))
}

type TrieNode struct {
	children map[string]*TrieNode
	isEnd    bool
	folder   string
}

func newNode() *TrieNode {
	return &TrieNode{
		children: make(map[string]*TrieNode),
		isEnd:    false,
	}
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) Insert(paths []string, s string) {
	curr := t.root
	for _, path := range paths {
		if curr.isEnd {
			return
		}
		if _, ok := curr.children[path]; !ok {
			curr.children[path] = newNode()
		}
		curr = curr.children[path]
	}
	curr.children = make(map[string]*TrieNode)
	curr.isEnd = true
	curr.folder = s
}

func removeSubfolders(folder []string) []string {
	var (
		trie Trie = Trie{root: newNode()}
		res  []string
		dfs  func(node *TrieNode)
	)
	for _, i := range folder {
		trie.Insert(split(i), i)
	}

	dfs = func(node *TrieNode) {
		if node.isEnd {
			res = append(res, node.folder)
			return
		}
		for _, v := range node.children {
			dfs(v)
		}
	}

	dfs(trie.root)

	return res
}