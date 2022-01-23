package shortestPathApp

import (
	"green/models/maze"
	"testing"
)

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

func TestBFS(t *testing.T) {
	tests := []struct {
		expected maze.Maze
		in       int
	}{
		{
			maze.Maze{Maze: [][]int32{
				{1, 1, 1, 1, 0, 1, 1, 1},
				{1, 0, 0, 0, 0, 0, 0, 1},
				{1, 0, 1, 1, 1, 0, 1, 1},
				{1, 0, 0, 0, 1, 0, 0, 1},
				{1, 1, 1, 0, 1, 1, 0, 1},
				{1, 0, 0, 0, 1, 0, 0, 1},
				{1, 0, 1, 1, 1, 0, 1, 1},
				{1, 0, 0, 2, 0, 0, 0, 1},
				{1, 1, 1, 1, 1, 1, 1, 1},
			}},
			12,
		},
		{maze.Maze{Maze: [][]int32{
			{1, 1, 1, 1, 1, 1, 1, 1},
			{1, 0, 0, 0, 0, 0, 0, 1},
			{1, 0, 1, 1, 1, 0, 1, 1},
			{0, 0, 0, 0, 1, 0, 0, 1},
			{1, 0, 1, 0, 1, 1, 0, 1},
			{1, 0, 1, 1, 1, 0, 0, 1},
			{1, 0, 0, 2, 1, 0, 0, 1},
			{1, 1, 1, 1, 1, 1, 1, 1},
		}},
			6,
		},
		{maze.Maze{Maze: [][]int32{
			{1, 1, 1, 1, 1, 1, 1, 1},
			{1, 0, 0, 0, 0, 0, 0, 1},
			{1, 0, 1, 1, 1, 1, 0, 1},
			{1, 0, 0, 0, 1, 1, 0, 0},
			{1, 0, 1, 0, 1, 1, 1, 1},
			{1, 0, 1, 1, 1, 0, 0, 1},
			{1, 0, 0, 2, 1, 0, 0, 1},
			{1, 1, 1, 1, 1, 1, 1, 1},
		}},
			15,
		},
	}
	for _, test := range tests {
		destCordinate1, destCordinate2 := test.expected.GetDestinationCoordinates()
		sourceCoordinate1, sourceCoordinate2 := test.expected.GetSourceCoordinates()
		if BFS(test.expected.Maze, Point{sourceCoordinate1, sourceCoordinate2}, Point{destCordinate1, destCordinate2}) != test.in {
			t.Fail()
		}
	}
}
func TestGetShortestPath(t *testing.T) {
	tests := []struct {
		expected maze.Maze
		in       int
	}{
		{
			maze.Maze{Maze: [][]int32{
				{1, 1, 1, 1, 0, 1, 1, 1},
				{1, 0, 0, 0, 0, 0, 0, 1},
				{1, 0, 1, 1, 1, 0, 1, 1},
				{1, 0, 0, 0, 1, 0, 0, 1},
				{1, 1, 1, 0, 1, 1, 0, 1},
				{1, 0, 0, 0, 1, 0, 0, 1},
				{1, 0, 1, 1, 1, 0, 1, 1},
				{1, 0, 0, 2, 0, 0, 0, 1},
				{1, 1, 1, 1, 1, 1, 1, 1},
			}},
			12,
		},
		{maze.Maze{Maze: [][]int32{
			{1, 1, 1, 1, 1, 1, 1, 1},
			{1, 0, 0, 0, 0, 0, 0, 1},
			{1, 0, 1, 1, 1, 0, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 1, 0, 1, 1, 0, 1},
			{1, 0, 1, 1, 1, 0, 0, 1},
			{1, 0, 0, 2, 1, 0, 0, 1},
			{1, 1, 1, 1, 1, 1, 1, 1},
		}},
			5,
		},
		{maze.Maze{Maze: [][]int32{
			{1, 0, 1, 1, 1, 1, 1, 1},
			{1, 0, 1, 0, 0, 0, 0, 1},
			{1, 0, 1, 1, 1, 1, 0, 1},
			{1, 0, 0, 0, 1, 1, 0, 0},
			{1, 0, 1, 0, 1, 1, 1, 1},
			{1, 0, 1, 1, 1, 0, 0, 1},
			{1, 0, 0, 2, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1},
		}},
			8,
		},
	}
	for _, test := range tests {
		if GetShortestPath(&test.expected) != test.in {
			t.Errorf("expected: %d, actual: %d", test.in, GetShortestPath(&test.expected))
		}
	}
}

func TestQueue_BuildQueue(t *testing.T) {
	q.BuildQueue()
	if q.capacity != 1 {
		t.Errorf("expected: 1, actual: %d", q.capacity)
	}
	if q.size != 0 {
		t.Errorf("expected: 1, actual: %d", q.size)
	}
	if q.head != 0 {
		t.Errorf("expected: 1, actual: %d", q.head)
	}
	if len(q.queue) != 1 {
		t.Errorf("expected: 1, actual: %d", len(q.queue))
	}
}

func TestQueue_Front(t *testing.T) {
	q.BuildQueue()
	q.Push(&QueueNode{Point{0, 0}, 0})
	q.Push(&QueueNode{Point{1, 1}, 0})
	q.Push(&QueueNode{})

	if q.Front() != q.Pop() {
		t.Errorf("expected: %v, actual: %v", q.Pop(), q.Front())
	}
}
