package mazeService

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"green/apps/shortestPathApp"
	"green/database"
	"green/models/maze"
	"log"
)

const (
	DELETED               string = "deleted."
	INSERTED                     = "inserted."
	UPDATED                      = "updated."
	CollectionName               = "maze"
	NotRectangular               = "Maps must be rectangular. Please check it."
	InvalidMazeDimension         = "Maps cannot be larger than 100 in any dimension. Please check it."
	InvalidMazeSpaceValue        = "Map spaces can not use values other than the numbers 0-2 above. Please check it."
)

func InsertOne(newMaze *maze.Maze) (string, error) {
	switch {
	case !newMaze.IsRectangular():
		return NotRectangular, nil
	case newMaze.GetYDimension() > 100 || newMaze.GetXDimension() > 100:
		return InvalidMazeDimension, nil
	case false:
		return InvalidMazeSpaceValue, nil
	}

	isInserted, err := database.InsertOne(CollectionName,
		bson.M{"maze": newMaze.Maze, "mazeId": newMaze.MazeId, "shortestPath": shortestPathApp.GetShortestPath(newMaze)})
	if err != nil {
		log.Println("err in insert: ", err)
		return "", err
	}
	respBody := getRespBodyMessage(isInserted, INSERTED)
	return respBody, err
}

func DeleteOne(mazeId int) (string, error) {
	isDeleted, err := database.DeleteOne(CollectionName, bson.M{"mazeId": mazeId})
	if err != nil {
		log.Println("err in deletion: ", err)
		return "", err
	}
	respBody := getRespBodyMessage(isDeleted, DELETED)
	return respBody, err
}

func UpdateOne(updatedMaze *maze.Maze, mazeId int) (string, error) {
	isUpdated, err := database.UpdateOne(CollectionName, bson.M{"mazeId": mazeId}, bson.M{"maze": updatedMaze.Maze})
	if err != nil {
		log.Println("err in update: ", err)
		return "", err
	}
	respBody := getRespBodyMessage(isUpdated, UPDATED)
	return respBody, err
}

func FindById(maze *maze.Maze, mazeId int) error {
	err := database.FindById(CollectionName, bson.M{"mazeId": mazeId}, maze)
	if err != nil {
		log.Println("err in get: ", err)
		return err
	}
	return nil
}

func getRespBodyMessage(isSuccess bool, message string) string {
	messageBody := fmt.Sprintf("Couldn't %s", message)
	if isSuccess {
		messageBody = fmt.Sprintf("Successfully %s", message)
	}
	return messageBody
}
