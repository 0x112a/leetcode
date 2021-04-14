package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type Trie struct {
	child [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	n := this
	for _, v := range word {
		v -= 'a'
		if n.child[v] == nil {
			n.child[v] = &Trie{}
		}
		n = n.child[v]
	}
	n.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) SearchPrefix(word string) *Trie {
	n := this
	for _, v := range word {
		v -= 'a'
		if n.child[v] == nil {
			return nil
		}
		n = n.child[v]
	}
	return n
}

func (this *Trie) Search(word string) bool {
	n := this.SearchPrefix(word)
	return n != nil && n.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	n := this.SearchPrefix(prefix)
	return n != nil
}
