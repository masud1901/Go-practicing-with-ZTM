package queue


type queue struct {
	items []int
	capacity int
}

func new(capacity int) queue  {
	return queue{
		items:    make([]int, 0, capacity),
		capacity: capacity,
	}
}

func (q *queue) Enqueue(item int) bool {
	if len(q.items) >= q.capacity {
		return false
	}
	
	q.items = append(q.items, item)
	return true
}

func (q *queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *queue) Size() int {
	return len(q.items)
}