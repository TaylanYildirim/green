package mazeService

import (
	"go.mongodb.org/mongo-driver/bson"
	"green/apps/shortestPathApp"
	"green/database"
	"green/models/maze"
)

func InsertOne(newMaze *maze.Maze) (bool, error) {
	isInserted, err := database.InsertOne("maze",
		bson.M{"maze": newMaze.Maze, "mazeId": newMaze.MazeId, "shortestPath": shortestPathApp.GetShortestPath(newMaze)})
	if err != nil {
		return false, err
	}
	return isInserted, err
}

func DeleteOne(mazeId int) (bool, error) {
	isDeleted, err := database.DeleteOne("maze", bson.M{"mazeId": mazeId})
	if err != nil {
		return false, err
	}
	return isDeleted, err
}

func UpdateOne(updatedMaze *maze.Maze, mazeId int) (bool, error) {
	isUpdated, err := database.UpdateOne("maze", bson.M{"mazeId": mazeId}, bson.M{"maze": updatedMaze.Maze})
	if err != nil {
		return false, err
	}
	return isUpdated, err
}

func FindById(maze *maze.Maze, mazeId int) error {
	err := database.FindById("maze", bson.M{"mazeId": mazeId}, maze)
	if err != nil {
		return err
	}
	return nil
}
