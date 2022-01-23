package shortestPathApp

import "testing"

var q Queue

func init() {
	q.BuildQueue()
}
func TestQueue_Push(t *testing.T) {
	q.Push(&QueueNode{})
	q.Push(&QueueNode{})
	q.Push(&QueueNode{})
	q.Push(&QueueNode{})

	if size := len(q.queue); size != 4 {
		t.Errorf("wrong count, expected 4 and got %d", size)
	}
}
func TestQueue_Pop(t *testing.T) {
	q.BuildQueue()
	q.Push(&QueueNode{})
	q.Push(&QueueNode{})
	q.Pop()
	q.Pop()

	if q.head != 0 {
		t.Errorf("wrong count, expected 0 and got %d", q.head)
	}
}
func TestQueue_IsFull(t *testing.T) {
	q.BuildQueue()
	q.size = 2
	q.capacity = 2
	if !q.IsFull() {
		t.Errorf("IsFull should return true")
	}
}
func TestQueue_IsEmpty(t *testing.T) {
	q.BuildQueue()
	q.Push(&QueueNode{})
	q.Push(&QueueNode{})
	q.Pop()
	q.Pop()
	if !q.IsEmpty() {
		t.Errorf("IsEmpty should return true")
	}
}
