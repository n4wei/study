package heap

// 4:28-5:08

type Comparator func(a, b int) bool

type Heap struct {
	store []int
	compareFunc Comparator
}

func NewHeap(nums []int, compareFunc Comparator) *Heap {
	heap := &Heap{
		store: nums,
		compareFunc: compareFunc,
	}
	if len(heap.store) > 0 {
		heap.buildHeap()
	}
	return heap
}

// runtime O(n)
// space O(n)
func (this *Heap) buildHeap() {
	i := len(this.store)/2
	for i >= 0 {
		this.heapify(i)
		i--
	}
}

// runtime O(logn)
func (this *Heap) heapify(i int) {
	var parentIndex, leftChildIndex, rightChildIndex int
	for {
		parentIndex = i
		leftChildIndex = 2*i+1
		rightChildIndex = leftChildIndex+1

		if leftChildIndex < len(this.store) && this.compareFunc(this.store[leftChildIndex], this.store[parentIndex]) {
			parentIndex = leftChildIndex
		}
		if rightChildIndex < len(this.store) && this.compareFunc(this.store[rightChildIndex], this.store[parentIndex]) {
			parentIndex = rightChildIndex
		}

		if parentIndex == i {
			return
		}
		this.store[i], this.store[parentIndex] = this.store[parentIndex], this.store[i]
		i = parentIndex
	}
}

// runtime O(logn)
func (this *Heap) Push(n int) {
	this.store = append(this.store, n)
	childIndex := len(this.store)-1
	var parentIndex int
	
	for {
		parentIndex = (childIndex-1)/2
		if parentIndex >= 0 && this.compareFunc(this.store[childIndex], this.store[parentIndex]) {
			this.store[childIndex], this.store[parentIndex] = this.store[parentIndex], this.store[childIndex]
			childIndex = parentIndex
		} else {
			return
		}
	}
}

// runtime O(1)
func (this *Heap) Top() int {
	if len(this.store) == 0 {
		return -1
	}
	return this.store[0]
}

// runtime O(logn)
func (this *Heap) Pop() {
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
func (this *Heap) Len() int {
	return len(this.store)
}