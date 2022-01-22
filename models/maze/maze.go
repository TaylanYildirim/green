package maze

type Maze struct {
	Maze   [][]int32 `bson:"maze,omitempty"`
	MazeId int32     `bson:"mazeId,omitempty"`
}
