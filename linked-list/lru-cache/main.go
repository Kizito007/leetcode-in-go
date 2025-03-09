package main

import (
	"fmt"
)

// Node represents an entry in the cache.
type Node struct {
	key, value int
	prev, next *Node
}

// LRUCache is the main structure that supports get and put operations.
type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node // Dummy head.
	tail     *Node // Dummy tail.
}

// Constructor initializes the LRUCache with a given capacity.
func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

// Get returns the value associated with the key if present, else -1.
func (this *LRUCache) Get(key int) int {
	if node, exists := this.cache[key]; exists {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

// Put inserts a key-value pair into the cache. If the key exists, update its value.
// If adding a new key causes the cache to exceed capacity, evict the least recently used item.
func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.cache[key]; exists {
		node.value = value
		this.moveToHead(node)
	} else {
		newNode := &Node{
			key:   key,
			value: value,
		}
		this.cache[key] = newNode
		this.addToHead(newNode)
		if len(this.cache) > this.capacity {
			tailPrev := this.tail.prev
			this.removeNode(tailPrev)
			delete(this.cache, tailPrev.key)
		}
	}
}

// removeNode removes a node from the linked list.
func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// addToHead inserts a node right after the dummy head.
func (this *LRUCache) addToHead(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

// moveToHead moves an existing node to the head (marking it as most recently used).
func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1)) // prints 1, making key 1 the most recently used
	lru.Put(3, 3)           // evicts key 2 because capacity is exceeded
	fmt.Println(lru.Get(2)) // prints -1 (key 2 was evicted)
	lru.Put(4, 4)           // evicts key 1
	fmt.Println(lru.Get(1)) // prints -1 (key 1 was evicted)
	fmt.Println(lru.Get(3)) // prints 3
	fmt.Println(lru.Get(4)) // prints 4
}
