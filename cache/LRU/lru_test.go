package LRU

import (
	"errors"
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	type fields struct {
		mapping  map[string]*Node
		cache    *DoubleLinkedList
		capacity int
	}
	type args struct {
		key string
	}

	t.Run("test LRU cache", func(t *testing.T) {
		l := NewLRUCache(5)

		_, err := l.Get("hello")
		if err == nil {
			t.Errorf("Get(hello) should return error")
			return
		} else if !errors.Is(err, keyNotFound) {
			t.Errorf("Get(hello) should return error %v, but got %v", keyNotFound, err)
			return
		}
		l.Put("hello", "world")
		l.Put("1", "David")
		l.Put("2", "Jack")
		l.Put("3", "Will")
		l.Put("4", "Andy")
		l.Put("5", "Steve")
		lookUp(l.cache)

		_, err = l.Get("hello")
		if err == nil {
			t.Errorf("Get(hello) should return error")
			return
		} else if !errors.Is(err, keyNotFound) {
			t.Errorf("Get(hello) should return error %v, but got %v", keyNotFound, err)
			return
		}

	})
}

func lookUp(cache *DoubleLinkedList) {
	first := cache.head.next
	for first != cache.tail {
		fmt.Printf("node = %+v\n", first)
		first = first.next
	}
}
