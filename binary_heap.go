package main

type LessFunc func(a interface{}, b interface{}) bool

type MaxBinaryHeap struct {
	values []interface{}
	isLess LessFunc
}

func (heap *MaxBinaryHeap) swapValues(a int, b int) {
	heap.values[a], heap.values[b] = heap.values[b], heap.values[a]
}

func (heap *MaxBinaryHeap) bubbleUp(index int) {
	if index == 0 {
		return
	}

	parentIndex := (index - 1) / 2

	if heap.isLess(heap.values[parentIndex], heap.values[index]) {
		heap.swapValues(parentIndex, index)
		heap.bubbleUp(parentIndex)
	}
}

func (heap *MaxBinaryHeap) sinkDown(index int) {
	lastIndex := len(heap.values) - 1
	leftChildIndex := (index * 2) + 1
	rightChildIndex := (index * 2) + 2

	if index >= lastIndex {
		return
	}

	if leftChildIndex <= lastIndex &&
		heap.isLess(heap.values[index], heap.values[leftChildIndex]) {
		heap.swapValues(index, leftChildIndex)
		heap.sinkDown(leftChildIndex)
	}

	if rightChildIndex <= lastIndex &&
		heap.isLess(heap.values[index], heap.values[rightChildIndex]) &&
		heap.isLess(heap.values[leftChildIndex], heap.values[rightChildIndex]) {
		heap.swapValues(index, rightChildIndex)
		heap.sinkDown(rightChildIndex)
	}
}

func (heap *MaxBinaryHeap) Insert(val interface{}) {
	heap.values = append(heap.values, val)
	heap.bubbleUp(len(heap.values) - 1)
}

func (heap *MaxBinaryHeap) ExtractMax() interface{} {
	lastIndex := len(heap.values) - 1
	if lastIndex < 0 {
		return nil
	}

	heap.swapValues(0, lastIndex)

	removed := heap.values[lastIndex]
	heap.values = heap.values[:lastIndex]

	heap.sinkDown(0)

	return removed
}

func NewMaxBinaryHeap(lessFunc LessFunc) MaxBinaryHeap {
	return MaxBinaryHeap{
		isLess: lessFunc,
	}
}

func NewPriorityQueue(lessFunc LessFunc) MaxBinaryHeap {
	return MaxBinaryHeap{
		isLess: lessFunc,
	}
}
