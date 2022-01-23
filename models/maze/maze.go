package maze

type Maze struct {
	Maze         [][]int32 `bson:"maze,omitempty"`
	MazeId       int32     `bson:"mazeId,omitempty"`
	ShortestPath int32     `bson:"shortestPath"`
}

func (M *Maze) GetYDimension() int {
	return len(M.Maze[0])
}

func (M *Maze) GetXDimension() int {
	return len(M.Maze)
}

func (M *Maze) GetSourceCoordinates() (int, int) {
	for i := 0; i < M.GetXDimension(); i++ {
		for j := 0; j < M.GetYDimension(); j++ {
			if M.Maze[i][j] == 2 {
				return i, j
			}
		}
	}
	return -1, -1
}

func (M *Maze) GetDestCoordinates() (int, int) {
	for i := 0; i < M.GetXDimension(); i++ {
		for j := 0; j < M.GetYDimension(); j++ {
			switch {
			case i == 0 && M.Maze[i][j] == 0,
				i == M.GetXDimension()-1 && M.Maze[i][j] == 0,
				j == 0 && M.Maze[i][j] == 0,
				j == M.GetYDimension()-1 && M.Maze[i][j] == 0:
				return i, j
			}
		}
	}
	return -1, -1
}
