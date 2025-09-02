package priorityqueue

import (
	"testing"
)

func TestPriorityQueueInt(t *testing.T) {
	// Test with integer priorities
	pq := New(func(item string) int {
		// Priority based on string length
		return len(item)
	})

	// Enqueue items
	pq.Enqueue("short")
	pq.Enqueue("medium_length")
	pq.Enqueue("very_long_item")

	// Test Peek (should return highest priority - longest string)
	item, ok := pq.Peek()
	if !ok || item != "very_long_item" {
		t.Errorf("Expected 'very_long_item' from Peek, got %v", item)
	}

	// Test Dequeue order (highest priority first - longest strings first)
	expected := []string{"very_long_item", "medium_length", "short"}
	for _, exp := range expected {
		item, ok := pq.Dequeue()
		if !ok || item != exp {
			t.Errorf("Expected %s from Dequeue, got %v", exp, item)
		}
	}

	// Test empty queue
	_, ok = pq.Dequeue()
	if ok {
		t.Error("Expected Dequeue to return false on empty queue")
	}
}

func TestPriorityQueueFloat(t *testing.T) {
	// Test with float priorities
	type Task struct {
		name     string
		priority float64
	}

	pq := New(func(task Task) float64 {
		return task.priority
	})

	// Enqueue tasks with different priorities
	pq.Enqueue(Task{name: "low", priority: 1.0})
	pq.Enqueue(Task{name: "high", priority: 10.0})
	pq.Enqueue(Task{name: "medium", priority: 5.0})

	// Test Dequeue order (should be high, medium, low)
	expected := []string{"high", "medium", "low"}
	for _, exp := range expected {
		task, ok := pq.Dequeue()
		if !ok || task.name != exp {
			t.Errorf("Expected task with name %s from Dequeue, got %v", exp, task)
		}
	}
}

func TestPriorityQueueString(t *testing.T) {
	// Test with string priorities
	pq := New(func(item int) string {
		// Priority based on string representation
		return string(rune(item + 64)) // A=65, B=66, etc.
	})

	// Enqueue items
	pq.Enqueue(3) // C (highest priority)
	pq.Enqueue(1) // A (lowest priority)
	pq.Enqueue(2) // B (medium priority)

	// Test Dequeue order (highest priority first - C, B, A)
	expected := []int{3, 2, 1}
	for _, exp := range expected {
		item, ok := pq.Dequeue()
		if !ok || item != exp {
			t.Errorf("Expected %d from Dequeue, got %v", exp, item)
		}
	}
}

func TestPriorityQueueEmpty(t *testing.T) {
	pq := New(func(item string) int {
		return len(item)
	})

	if !pq.IsEmpty() {
		t.Error("Expected new queue to be empty")
	}

	_, ok := pq.Peek()
	if ok {
		t.Error("Expected Peek to return false on empty queue")
	}

	_, ok = pq.Dequeue()
	if ok {
		t.Error("Expected Dequeue to return false on empty queue")
	}

	pq.Enqueue("test")
	if pq.IsEmpty() {
		t.Error("Expected queue to not be empty after enqueue")
	}

	pq.Clear()
	if !pq.IsEmpty() {
		t.Error("Expected queue to be empty after clear")
	}
}

func TestPriorityQueueSingleItem(t *testing.T) {
	pq := New(func(item string) int {
		return len(item)
	})

	pq.Enqueue("single")

	item, ok := pq.Peek()
	if !ok || item != "single" {
		t.Errorf("Expected 'single' from Peek, got %v", item)
	}

	item, ok = pq.Dequeue()
	if !ok || item != "single" {
		t.Errorf("Expected 'single' from Dequeue, got %v", item)
	}

	if !pq.IsEmpty() {
		t.Error("Expected queue to be empty after dequeue")
	}
}
