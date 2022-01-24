package mazeService

import (
	"go.mongodb.org/mongo-driver/bson"
	"green/apps/shortestPathApp"
	"green/database"
	"green/models/maze"
	"log"
)

const (
	CollectionName        = "maze"
	NotRectangular        = "Maps must be rectangular, operation to DB failed. Please check maze."
	InvalidMazeDimension  = "Maps cannot be larger than 100 in any dimension, operation to DB failed. Please check maze."
	InvalidMazeSpaceValue = "Map spaces can not use values other than the numbers 0-2 above. Please check maze."
	GetSuccessMessage     = "Successfully retrieved."
	GetFailMessage        = "Invalid Maze id, this maze id does not exist."
)

func validateMaze(maze *maze.Maze) string {
	switch {
	case !maze.IsRectangular():
		return NotRectangular
	case maze.GetYDimension() > 100 || maze.GetXDimension() > 100:
		return InvalidMazeDimension
	case !maze.IsValidMazeSpaceValues():
		return InvalidMazeSpaceValue
	}
	return ""
}
func InsertOne(newMaze *maze.Maze) (string, error) {
	isValidMaze := validateMaze(newMaze)
	if isValidMaze != "" {
		return isValidMaze, nil
	}
	respMessage, err := database.InsertOne(CollectionName,
		bson.M{"maze": newMaze.Maze, "mazeId": newMaze.MazeId, "shortestPath": shortestPathApp.GetShortestPath(newMaze)})
	if err != nil {
		log.Println("err in insert: ", err)
		return "", err
	}

	return respMessage, err
}

func DeleteOne(mazeId int) (string, error) {
	respMessage, err := database.DeleteOne(CollectionName, bson.M{"mazeId": mazeId})
	if err != nil {
		log.Println("err in deletion: ", err)
		return "", err
	}

	return respMessage, nil
}

func UpdateOne(updatedMaze *maze.Maze, mazeId int) (string, error) {
	isValidMaze := validateMaze(updatedMaze)
	if isValidMaze != "" {
		return isValidMaze, nil
	}

	respMessage, err := database.UpdateOne(CollectionName, bson.M{"mazeId": mazeId},
		bson.M{"$set": bson.M{"maze": updatedMaze.Maze, "shortestPath": shortestPathApp.GetShortestPath(updatedMaze)}})

	if err != nil {
		log.Println("err in update: ", err)
		return respMessage, err
	}

	return respMessage, err
}

func FindById(maze *maze.Maze, mazeId int) (string, error) {
	err := database.FindById(CollectionName, bson.M{"mazeId": mazeId}, maze)
	if err != nil {
		log.Println("err in get: ", err)
		return GetFailMessage, err
	}
	return GetSuccessMessage, nil
}
