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
	NotRectangular               = "Maps must be rectangular, operation to DB failed. Please check maze."
	InvalidMazeDimension         = "Maps cannot be larger than 100 in any dimension, operation to DB failed. Please check maze."
	InvalidMazeSpaceValue        = "Map spaces can not use values other than the numbers 0-2 above. Please check maze."
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
	isValidMaze := validateMaze(updatedMaze)
	if isValidMaze != "" {
		return isValidMaze, nil
	}
	isUpdated, err := database.UpdateOne(CollectionName, bson.M{"mazeId": mazeId}, bson.M{"$set": bson.M{"maze": updatedMaze.Maze}})
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
