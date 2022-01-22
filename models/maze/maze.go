package maze

type Maze struct {
	Maze [][]uint8 `json:"maze"`
}

type MongoMaze struct {
	Maze   [][]int32 `bson:"maze,omitempty"`
	MazeId int32     `bson:"mazeId,omitempty"`
}
