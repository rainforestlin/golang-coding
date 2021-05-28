package main

import "fmt"

type LinkedNode struct {
	key, value int
	prev, next *LinkedNode
}

type LRUCache struct {
	capacity   int
	cache      map[int]*LinkedNode
	head, tail *LinkedNode
}

func initLinkedNode(key, value int) *LinkedNode {
	return &LinkedNode{
		value: value,
		key:   key,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    make(map[int]*LinkedNode),
		capacity: capacity,
		head:     initLinkedNode(0, 0),
		tail:     initLinkedNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *LRUCache) Get(key int) int {
	if _, ok := l.cache[key]; !ok {
		return -1
	}
	node := l.cache[key]
	l.moveToHead(node)
	return node.value
}

func (l *LRUCache) Put(key, value int) {
	if _, ok := l.cache[key]; !ok {
		node := initLinkedNode(key, value)
		l.cache[key] = node
		l.addToHead(node)
		if len(l.cache) > l.capacity {
			delete(l.cache, l.tail.prev.key)
			l.removeNode(l.tail.prev)
		}
	} else {
		node := l.cache[key]
		node.value = value
		l.moveToHead(node)
	}
}

func (l *LRUCache) addToHead(node *LinkedNode) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node

}

func (l *LRUCache) removeNode(node *LinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev

}

func (l *LRUCache) moveToHead(node *LinkedNode) {
	l.removeNode(node)
	l.addToHead(node)
}

func main() {
	lru := Constructor(3)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	fmt.Println(lru.Get(3))
	fmt.Println(lru.head.next.next.next.value)
}
