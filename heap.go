package aoc

// Heap is a priority queue implementation.
type Heap[T any] struct {
	items      []T
	comparator func(a, b T) bool
}

// NewHeap creates a new Heap using a comparator.
func NewHeap[T any](comparator func(a, b T) bool) Heap[T] {
	return Heap[T]{comparator: comparator}
}

// Push pushes a new item.
func (pq *Heap[T]) Push(item T) {
	pq.items = append(pq.items, item)
	pq.heapifyUp(len(pq.items) - 1)
}

// Pop removes and returns the top element from the priority queue.
func (pq *Heap[T]) Pop() T {
	top := pq.items[0]
	lastIndex := len(pq.items) - 1
	pq.items[0], pq.items[lastIndex] = pq.items[lastIndex], pq.items[0]
	pq.items = pq.items[:lastIndex]
	pq.heapifyDown(0)
	return top
}

// Peek returns the top element of the priority queue without removing it.
func (pq *Heap[T]) Peek() T {
	return pq.items[0]
}

// Len returns the number of elements in the priority queue.
func (pq *Heap[T]) Len() int {
	return len(pq.items)
}

// IsEmpty checks if the priority queue is empty.
func (pq *Heap[T]) IsEmpty() bool {
	return len(pq.items) == 0
}

func (pq *Heap[T]) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if pq.comparator(pq.items[index], pq.items[parentIndex]) {
			pq.items[index], pq.items[parentIndex] = pq.items[parentIndex], pq.items[index]
			index = parentIndex
		} else {
			break
		}
	}
}

func (pq *Heap[T]) heapifyDown(index int) {
	for {
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		smallest := index

		if leftChild < len(pq.items) && pq.comparator(pq.items[leftChild], pq.items[smallest]) {
			smallest = leftChild
		}

		if rightChild < len(pq.items) && pq.comparator(pq.items[rightChild], pq.items[smallest]) {
			smallest = rightChild
		}

		if smallest != index {
			pq.items[index], pq.items[smallest] = pq.items[smallest], pq.items[index]
			index = smallest
		} else {
			break
		}
	}
}
