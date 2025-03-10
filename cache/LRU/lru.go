package LRU

import "errors"

var (
	keyNotFound = errors.New("key not found")
)

type Node struct {
	key, val   string
	prev, next *Node
}
type DoubleLinkedList struct {
	head, tail *Node
	size       int
}
type LRUCache struct {
	mapping  map[string]*Node
	cache    *DoubleLinkedList
	capacity int
}

func NewDoubleLinkedList() *DoubleLinkedList {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &DoubleLinkedList{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (l *DoubleLinkedList) Size() int {
	return l.size
}
func (l *DoubleLinkedList) AddLast(node *Node) {
	node.prev = l.tail.prev
	node.next = l.tail
	l.tail.prev.next = node
	l.tail.prev = node
	l.size++
}
func (l *DoubleLinkedList) Remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	l.size--
}
func (l *DoubleLinkedList) RemoveFirst() *Node {
	if l.size == 0 || l.head.next == l.tail {
		return nil
	}
	first := l.head.next
	l.Remove(first)
	return first
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		mapping:  make(map[string]*Node),
		cache:    NewDoubleLinkedList(),
		capacity: capacity,
	}
}

func (l *LRUCache) Get(key string) (string, error) {
	n, ok := l.mapping[key]
	if !ok {
		return "", keyNotFound
	}

	// put this key to the tail
	l.cache.Remove(n)
	l.cache.AddLast(n)

	return n.val, nil
}

func (l *LRUCache) addRecently(key, val string) {
	// add to the tail
	n := &Node{key: key, val: val}
	l.cache.AddLast(n)
	l.mapping[key] = n
}
func (l *LRUCache) Put(key string, val string) {
	if n, ok := l.mapping[key]; ok {
		// delete key
		l.cache.Remove(n)
		delete(l.mapping, key)

		// add to the tail
		l.addRecently(key, val)
		return
	}

	// remove least access
	if l.capacity <= l.cache.Size() {
		removed := l.cache.RemoveFirst()
		if removed != nil {
			delete(l.mapping, removed.key)
		}
	}

	l.addRecently(key, val)
}
