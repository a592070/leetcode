package main

import "fmt"

func main() {
	trie := Constructor()
	trie.Insert("apple")
	//trie.Insert("app")
	//trie.Insert("butter")
	trie.Insert("butterfly")

	fmt.Printf("%+v\n", trie)
	fmt.Printf("trie.StartsWith(\"app\")=%v\n", trie.StartsWith("app"))
	fmt.Printf("trie.Search(\"apple\")=%v\n", trie.Search("apple"))
	fmt.Printf("trie.Search(\"app\")=%v\n", trie.Search("app"))
	fmt.Printf("trie.StartsWith(\"butter\")=%v\n", trie.StartsWith("butter"))
	fmt.Printf("trie.Search(\"butter\")=%v\n", trie.Search("butter"))
	fmt.Printf("trie.Search(\"butterfly\")=%v\n", trie.Search("butterfly"))
}

type Trie struct {
	head *node
}

type node struct {
	ch    int32
	isEnd bool
	next  []*node
}

func Constructor() Trie {
	return Trie{
		head: &node{ch: -1, isEnd: true, next: make([]*node, 26)},
	}
}

func (this *Trie) Insert(word string) {
	head := this.head
	for _, ch := range word {
		//fmt.Printf("%c->%d, %d\n", ch, ch, ch-97)
		if head.next[ch-97] == nil {
			n := &node{ch: ch, isEnd: false, next: make([]*node, 26)}
			head.next[ch-97] = n
		}
		head = head.next[ch-97]
	}
	head.isEnd = true
}

func (this *Trie) Search(word string) bool {
	head := this.head
	if len(head.next) == 0 {
		return false
	}
	for _, ch := range word {
		head = head.next[ch-97]
		if head == nil {
			return false
		}
		if head.ch != ch {
			return false
		}
	}
	return head.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	head := this.head
	if len(head.next) == 0 {
		return false
	}
	for _, ch := range prefix {
		head = head.next[ch-97]
		if head == nil {
			return false
		}
		if head.ch != ch {
			return false
		}
	}
	return true
}
