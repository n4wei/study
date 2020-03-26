package main

import (
	"fmt"
)

// 4:28-5:08
func main() {
	fmt.Println("min heap")
	heap := NewMinHeap([]int{5,4,3,2,1})
	fmt.Println(heap.store)
	heap = NewMinHeap([]int{})
	heap.Push(5)
	heap.Push(4)
	heap.Push(3)
	heap.Push(2)
	heap.Push(1)
	fmt.Println(heap.store)
	fmt.Println(heap.Len())
	fmt.Println(heap.Top())
	heap.Pop()
	fmt.Println(heap.Top())
	heap.Pop()
	fmt.Println(heap.Top())
	heap.Pop()
	fmt.Println(heap.Top())
	heap.Pop()
	fmt.Println(heap.Top())
	heap.Pop()
	fmt.Println(heap.Top())
	heap.Pop()
	fmt.Println(heap.Len())
	fmt.Println(heap.Top())
}

type MinHeap struct {
	store []int
}

func NewMinHeap(nums []int) *MinHeap {
	heap := &MinHeap{store: nums}
	if len(heap.store) > 0 {
		heap.buildHeap()
	}
	return heap
}

// runtime O(n)
// space O(n)
func (this *MinHeap) buildHeap() {
	i := len(this.store)/2
	for i >= 0 {
		this.heapify(i)
		i--
	}
}

// runtime O(logn)
func (this *MinHeap) heapify(i int) {
	var minIndex, leftChildIndex, rightChildIndex int
	for {
		minIndex = i
		leftChildIndex = 2*i+1
		rightChildIndex = leftChildIndex+1

		if leftChildIndex < len(this.store) && this.store[leftChildIndex] < this.store[minIndex] {
			minIndex = leftChildIndex
		}
		if rightChildIndex < len(this.store) && this.store[rightChildIndex] < this.store[minIndex] {
			minIndex = rightChildIndex
		}

		if minIndex == i {
			return
		}
		this.store[i], this.store[minIndex] = this.store[minIndex], this.store[i]
		i = minIndex
	}
}

// runtime O(logn)
func (this *MinHeap) Push(n int) {
	this.store = append(this.store, n)
	childIndex := len(this.store)-1
	var parentIndex int
	
	for {
		parentIndex = (childIndex-1)/2
		if parentIndex >= 0 && this.store[childIndex] < this.store[parentIndex] {
			this.store[childIndex], this.store[parentIndex] = this.store[parentIndex], this.store[childIndex]
			childIndex = parentIndex
		} else {
			return
		}
	}
}

// runtime O(1)
func (this *MinHeap) Top() int {
	if len(this.store) == 0 {
		return -1
	}
	return this.store[0]
}

// runtime O(logn)
func (this *MinHeap) Pop() {
	if len(this.store) == 0 {
		return
	}

	l := len(this.store)
	if l > 1 {
		this.store[0], this.store[l-1] = this.store[l-1], this.store[0]
	}
	this.store = this.store[:l-1]
	this.heapify(0)
}

// runtime O(1)
func (this *MinHeap) Len() int {
	return len(this.store)
}