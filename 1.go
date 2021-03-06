package main

const (
	hostbit = uint64(^uint(0)) == ^uint64(0)
	LENGTH  = 100
)

type lruNode struct {
	prev  *lruNode
	next  *lruNode
	key   int
	value int
	hnext *lruNode
}

type lruCache struct {
	node     []lruNode
	head     *lruNode
	tail     *lruNode
	capacity int
	used     int
}

func Construct(capacity int) lruCache {
	return lruCache{
		node:     make([]lruNode, LENGTH),
		head:     nil,
		tail:     nil,
		capacity: capacity,
		used:     0,
	}
}

func (this *lruCache) Get(key int) int {
	if this.tail == nil {
		return -1
	}
	if tmp := this.searchNode(key); tmp != nil {
		this.movetoTail(tmp)
		return tmp.value
	}
	return -1
}

func (this *lruCache) Put(key int, value int) {
	if tmp := this.searchNode(key); tmp != nil {
		tmp.value = value
		this.movetoTail(tmp)
		return
	}
	this.addNode(key, value)
	if this.used > this.capacity {
		this.delNode()
	}
}

func (this *lruCache) addNode(key int, value int) {
	newNode := &lruNode{
		key:   key,
		value: value,
	}
	tmp := &this.node[hash(key)]
	newNode.hnext = tmp.hnext
	tmp.hnext = newNode
	this.used++
	if this.tail == nil {
		this.tail, this.head = newNode, newNode
		return
	}
	this.tail.next = newNode
	newNode.prev = this.tail
	this.tail = newNode
}

func (this *lruCache) delNode() {
	if this.head == nil {
		return
	}
	prev := &this.node[hash(this.head.key)]
	tmp := prev.next
	for tmp != nil && tmp.key != this.head.key {
		prev = tmp
		tmp = tmp.next
	}
	if tmp == nil {
		return
	}
	prev.hnext = tmp.hnext
	this.head = this.head.next
	this.head.prev = nil
	this.used--
}

func (this *lruCache) searchNode(key int) *lruNode {
	if this.tail == nil {
		return nil
	}
	tmp := this.node[hash(key)].hnext
	for tmp != nil {
		if tmp.key == key {
			return tmp
		}
		tmp = tmp.hnext
	}
	return nil
}

func (this *lruCache) movetoTail(node *lruNode) {
	if this.tail == node {
		return
	}
	if this.head == node {
		this.head = node.next
		this.head.prev = nil
	} else {
		node.next.prev = node.prev
		node.prev.next = node.next
	}
	node.next = nil
	this.tail.next = node
	node.prev = this.tail
	this.tail = node
}

func hash(key int) int {
	if hostbit {
		return (key ^ (key >> 32)) & (LENGTH - 1)
	}
	return (key ^ (key >> 16)) & (LENGTH - 1)
}
