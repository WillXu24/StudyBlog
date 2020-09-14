package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Insert("apple")
	fmt.Println(obj.Search("apple"))
	fmt.Println(obj.Search("app"))
	fmt.Println(obj.StartsWith("app"))

	obj.Insert("app")
	fmt.Println(obj.Search("app"))
}

// Trie 前缀树
type Trie struct {
	isEnd bool
	next  [26]*Trie
}

//Constructor Initialize your data structure here.
func Constructor() Trie {
	return Trie{}
}

//Insert Inserts a word into the trie.
func (this *Trie) Insert(word string) {
	for _, v := range word {
		if this.next[v-'a'] == nil {
			this.next[v-'a'] = new(Trie)
		}
		this = this.next[v-'a']
	}
	this.isEnd = true
}

// Search Returns if the word is in the trie.
func (this *Trie) Search(word string) bool {
	for _, v := range word {
		if this.next[v-'a'] == nil {
			return false
		}
		this = this.next[v-'a']
	}
	if this.isEnd {
		return true
	}
	return false
}

// StartsWith Returns if there is any word in the trie that starts with the given prefix.
func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if this.next[v-'a'] == nil {
			return false
		}
		this = this.next[v-'a']
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