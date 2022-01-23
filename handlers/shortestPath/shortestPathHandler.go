package shortestPath

import (
	"green/models/maze"
)

const (
	ROW int = 9
	COL     = 10
)

type Point struct {
	x int
	y int
}

type QueueNode struct {
	pt   Point
	dist int
}

func isValid(row int, col int) bool {
	return row >= 0 && row < ROW && col >= 0 && col < COL
}

func BFS(mat [][]int32, src Point, dest Point) int {
	rowNum := []int{-1, 0, 0, 1}
	colNum := []int{0, -1, 1, 0}
	if mat[src.x][src.y] == 1 || mat[dest.x][dest.y] == 1 {
		return -1
	}
	var visited [ROW][COL]bool
	visited[src.x][src.y] = true
	s := QueueNode{src, 0}

	var q Queue
	q.BuildQueue()
	q.Push(&s)

	for !q.IsEmpty() {
		curr := q.Front()
		pt := curr.pt
		if pt.x == dest.x && pt.y == dest.y {
			return curr.dist
		}
		q.Pop()
		for i := 0; i < 4; i++ {
			row := pt.x + rowNum[i]
			col := pt.y + colNum[i]
			if isValid(row, col) && mat[row][col] == 0 && !visited[row][col] {
				visited[row][col] = true
				Adcell := QueueNode{
					pt:   Point{row, col},
					dist: curr.dist + 1,
				}
				q.Push(&Adcell)
			}
		}

	}

	return -1
}

func GetShortestPath(newMaze *maze.Maze) int {
	sourceXCoordinate, sourceYCoordinate := newMaze.GetSourceCoordinates()
	destXCoordinate, destYCoordinate := newMaze.GetDestCoordinates()
	return BFS(newMaze.Maze, Point{sourceXCoordinate, sourceYCoordinate}, Point{destXCoordinate, destYCoordinate})
}

type Queue struct {
	queue                []*QueueNode
	size, head, capacity int
}

func (Q *Queue) BuildQueue() {
	Q.capacity = 1
	Q.queue = make([]*QueueNode, Q.capacity)
	Q.size = 0
	Q.head = 0
}

func (Q *Queue) Pop() (g *QueueNode) {
	if Q.IsEmpty() {
		panic("pop on empty queue!")
	}
	Q.size = Q.size - 1
	temp := Q.queue[Q.head]
	Q.head = (Q.head + 1) % Q.capacity
	return temp
}

func (Q *Queue) Front() *QueueNode {
	return Q.queue[Q.head]
}

func (Q *Queue) Push(g *QueueNode) {
	if Q.IsFull() {
		newSlice := make([]*QueueNode, Q.capacity*2)
		copy(newSlice, Q.queue)
		Q.queue = newSlice
		Q.capacity *= 2
	}
	Q.queue[(Q.head+Q.size)%Q.capacity] = g
	Q.size = Q.size + 1
}

func (Q *Queue) IsEmpty() bool {
	return Q.size == 0
}
func (Q *Queue) IsFull() bool {
	return Q.size == Q.capacity
}
