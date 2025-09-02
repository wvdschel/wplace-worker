package priorityqueue

import (
	"container/heap"
)

// Ordered is a constraint that matches any ordered type.
// An ordered type is one that supports the <, <=, >, and >= operators.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// PriorityQueue is a generic priority queue implementation
// that uses a user-provided priority function to determine element ordering.
type PriorityQueue[T any, P Ordered] struct {
	items    []T
	priority func(T) P
}

// New creates a new priority queue with the given priority function.
func New[T any, P Ordered](priorityFunc func(T) P) *PriorityQueue[T, P] {
	return &PriorityQueue[T, P]{
		items:    make([]T, 0),
		priority: priorityFunc,
	}
}

// Len returns the number of elements in the priority queue.
func (pq *PriorityQueue[T, P]) Len() int {
	return len(pq.items)
}

// Less reports whether the element with index i should sort before the element with index j.
// For a priority queue, we want higher priority items first, so we reverse the comparison.
func (pq *PriorityQueue[T, P]) Less(i, j int) bool {
	return pq.priority(pq.items[i]) > pq.priority(pq.items[j])
}

// Swap swaps the elements with indexes i and j.
func (pq *PriorityQueue[T, P]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

// Push adds an element to the priority queue.
func (pq *PriorityQueue[T, P]) Push(x interface{}) {
	item := x.(T)
	pq.items = append(pq.items, item)
}

// Pop removes and returns the highest priority element from the priority queue.
func (pq *PriorityQueue[T, P]) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	pq.items = old[0 : n-1]
	return item
}

// Enqueue adds an element to the priority queue.
func (pq *PriorityQueue[T, P]) Enqueue(item T) {
	heap.Push(pq, item)
}

// Dequeue removes and returns the highest priority element from the priority queue.
// Returns false if the queue is empty.
func (pq *PriorityQueue[T, P]) Dequeue() (T, bool) {
	if pq.Len() == 0 {
		var zero T
		return zero, false
	}
	item := heap.Pop(pq).(T)
	return item, true
}

// Peek returns the highest priority element without removing it from the queue.
// Returns false if the queue is empty.
func (pq *PriorityQueue[T, P]) Peek() (T, bool) {
	if pq.Len() == 0 {
		var zero T
		return zero, false
	}
	return pq.items[0], true
}

// IsEmpty returns true if the priority queue is empty.
func (pq *PriorityQueue[T, P]) IsEmpty() bool {
	return pq.Len() == 0
}

// Clear removes all elements from the priority queue.
func (pq *PriorityQueue[T, P]) Clear() {
	pq.items = make([]T, 0)
}
