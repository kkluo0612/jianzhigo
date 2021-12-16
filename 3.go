package main

import (
	"fmt"
)

type Node struct {
	value    int
	priority int
}

type Queue struct {
	heap     []Node
	capacity int
	used     int
}

func NewPriorityQueue(capacity int) Queue {
	return Queue{
		heap:     make([]Node, capacity+1, capacity+1),
		capacity: capacity,
		used:     0,
	}
}

func (q *Queue) Push(node Node) {
	if q.used >= q.capacity {
		return
	}
	q.used++
	q.heap[q.used] = node
}

func (q *Queue) Pop() Node {
	if q.used == 0 {
		return Node{-1, -1}
	}
	adjustNode(q.heap, 1, q.used)
	node := q.heap[1]
	q.heap[1] = q.heap[q.used]
	q.used--
	return node
}

func (q *Queue) Top() Node {
	if q.used == 0 {
		return Node{-1, -1}
	}
	adjustNode(q.heap, 1, q.used)
	return q.heap[1]
}

func adjustNode(heap []Node, start, end int) {
	if start >= end {
		return
	}
	for i := end / 2; i >= start; i-- {
		high := i
		if heap[high].priority < heap[2*i].priority {
			high = 2 * i
		}
		if 2*i+1 <= end && heap[high].priority < heap[2*i+1].priority {
			high = 2 * i
		}
		if high == i {
			continue
		}
		heap[high], heap[i] = heap[i], heap[high]
	}
}

func main() {
	queue := NewPriorityQueue(7)
	queue.Push(Node{0, 1})
	queue.Push(Node{3, 1})
	queue.Push(Node{3, 2})
	queue.Push(Node{6, 6})
	queue.Push(Node{12, 5})
	queue.Push(Node{13, 8})
	fmt.Println(queue)
	queue.Pop()
	fmt.Println(queue)
}
