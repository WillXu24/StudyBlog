package main

import "fmt"

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))

	lru.Put(3, 3)
	fmt.Println(lru.Get(2))

	lru.Put(4, 4)

	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(4))

}

/**********************************************************************************************************************/

type LRUCache struct {
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
	len, cap   int
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cache: map[int]*DLinkedNode{},
		head:  initDLinkedNode(0, 0),
		tail:  initDLinkedNode(0, 0),
		cap:   capacity,
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if ok {
		node.value = value
		this.moveToHead(node)
		return
	}
	newNode := initDLinkedNode(key, value)
	this.cache[key] = newNode
	this.addToHead(newNode)
	this.len++
	if this.len > this.cap {
		tailNode := this.removeTail()
		delete(this.cache, tailNode.key)
		this.len--
	}
}
