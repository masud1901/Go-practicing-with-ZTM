package queue

import "testing"

func TestQueue_EnqueueDequeue(t *testing.T) {
	q := new(2)

	if !q.Enqueue(1) {
		t.Errorf("Expected Enqueue to succeed")
	}
	if !q.Enqueue(2) {
		t.Errorf("Expected Enqueue to succeed")
	}
	if q.Enqueue(3) {
		t.Errorf("Expected Enqueue to fail due to capacity")
	}

	item, ok := q.Dequeue()
	if !ok || item != 1 {
		t.Errorf("Expected Dequeue to return 1, got %d", item)
	}

	item, ok = q.Dequeue()
	if !ok || item != 2 {
		t.Errorf("Expected Dequeue to return 2, got %d", item)
	}

	_, ok = q.Dequeue()
	if ok {
		t.Errorf("Expected Dequeue to fail on empty queue")
	}
}

func TestQueue_IsEmptyAndSize(t *testing.T) {
	q := new(3)

	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty")
	}
	if q.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", q.Size())
	}

	q.Enqueue(1)
	q.Enqueue(2)

	if q.IsEmpty() {
		t.Errorf("Expected queue to not be empty")
	}
	if q.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", q.Size())
	}

	q.Dequeue()
	if q.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", q.Size())
	}
}
