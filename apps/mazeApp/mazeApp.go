package mazeApp

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"green/models/maze"
	"green/services/mazeService"
	"green/utils/httpUtil"
	"green/utils/stringUtil"
	"io/ioutil"
	"log"
	"net/http"
)

func InsertMaze() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var newMaze maze.Maze
		jsonBody, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(jsonBody, &newMaze)
		if err != nil {
			log.Println("unmarshall err: ", err)
		}

		message, err := mazeService.InsertOne(&newMaze)

		httpUtil.GenerateResponse(w, r, err, map[string]interface{}{
			"message": message,
			"mazeId":  newMaze.MazeId,
		})
	}
	return fn
}

func DeleteMaze() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		mazeId := stringUtil.ParseUint(chi.URLParam(r, "id"))

		message, err := mazeService.DeleteOne(mazeId)

		httpUtil.GenerateResponse(w, r, err, map[string]interface{}{
			"message": message,
			"mazeId":  mazeId,
		})
	}
	return fn
}

func UpdateMaze() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var updatedMaze maze.Maze
		updatedMazeId := stringUtil.ParseUint(chi.URLParam(r, "id"))
		jsonBody, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(jsonBody, &updatedMaze)
		if err != nil {
			log.Println("unmarshal HTTP body err: ", err)
		}

		message, err := mazeService.UpdateOne(&updatedMaze, updatedMazeId)

		httpUtil.GenerateResponse(w, r, err, map[string]interface{}{
			"message": message,
			"mazeId":  updatedMazeId,
		})
	}
	return fn
}

func GetMaze() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var maze maze.Maze
		mazeId := stringUtil.ParseUint(chi.URLParam(r, "id"))

		message, _ := mazeService.FindById(&maze, mazeId)

		httpUtil.JSON(w, r, map[string]interface{}{
			"message": message,
			"mazeId":  mazeId,
			"maze":    maze,
		})
	}
	return fn
}
