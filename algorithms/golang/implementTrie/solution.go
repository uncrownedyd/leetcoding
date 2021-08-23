package implementTrie

type Trie struct {
	endOfWord bool
	next      [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	for _, char := range word {
		index := char - 'a'
		if t.next[index] == nil {
			t.next[index] = &Trie{}
		}
		t = t.next[index]
	}
	t.endOfWord = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	for _, char := range word {
		index := char - 'a'
		if t.next[index] == nil {
			return false
		}
		t = t.next[index]
	}
	return t.endOfWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	for _, char := range prefix {
		index := char - 'a'
		if t.next[index] == nil {
			return false
		}
		t = t.next[index]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
