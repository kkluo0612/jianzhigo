package main

import "fmt"

type Heap struct {
	a     []int
	n     int
	count int
}

//工厂方法
func NewHeap(capacity int) *Heap {
	heap := &Heap{}
	heap.n = capacity
	heap.a = make([]int, capacity+1)
	heap.count = 0
	return heap
}

func (heap *Heap) insert(data int) {
	if heap.count == heap.n {
		return
	}
	heap.count++
	heap.a[heap.count] = data
	i := heap.count
	parent := i / 2
	for parent > 0 && heap.a[parent] < heap.a[i] {
		swap(heap.a, parent, i)
		i = parent
		parent = i / 2
	}
}

func (heap *Heap) removeMax() {
	if heap.count == 0 {
		return
	}
	swap(heap.a, 1, heap.count)
	heap.count--
	uptodown(heap.a, heap.count)
}

func buildHeap(a []int, n int) {
	for i := n / 2; i >= 1; i-- {
		uptodown(a, i, n)
	}
}

func swap(a []int, i, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}

func uptodown(a []int, count int) {
	for i := 0; i <= count/2; {
		maxIndex := i
		if a[i] < a[i*2] {
			maxIndex = i * 2
		}
		if i*2+1 <= count && a[maxIndex] < a[i*2+1] {
			maxIndex = i*2 + 1
		}
		if maxIndex == i {
			break
		}
		swap(a, i, maxIndex)
		i = maxIndex
	}
}

func main() {
	heap := NewHeap(6)
	heap.insert(3)
	heap.insert(9)
	heap.insert(1)
	heap.insert(8)
	heap.insert(7)
	heap.insert(3)
	fmt.Println(heap)
	heap.removeMax()
	fmt.Println(heap)
}
