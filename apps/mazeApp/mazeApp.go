package mazeApp

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"green/models/maze"
	"green/services/mazeService"
	"green/utils/httpUtil"
	"green/utils/stringUtil"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	DELETED  string = "deleted."
	INSERTED        = "inserted."
	UPDATED         = "updated."
)

func InsertMaze() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var newMaze maze.Maze
		jsonBody, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(jsonBody, &newMaze)
		if err != nil {
			log.Println("unmarshall err: ", err)
		}

		isInserted, err := mazeService.InsertOne(&newMaze)

		httpUtil.GenerateResponse(w, r, err, map[string]interface{}{
			"message": getHTTPBodyMessage(w, isInserted, INSERTED),
			"mazeId":  newMaze.MazeId,
		})
	}
	return fn
}

func DeleteMaze() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		mazeId := stringUtil.ParseUint(chi.URLParam(r, "id"))
		isDeleted, err := mazeService.DeleteOne(mazeId)
		if err != nil {
			log.Println("err in deletion: ", err)
		}

		httpUtil.GenerateResponse(w, r, err, map[string]interface{}{
			"message": getHTTPBodyMessage(w, isDeleted, DELETED),
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

		isUpdated, err := mazeService.UpdateOne(&updatedMaze, updatedMazeId)

		httpUtil.GenerateResponse(w, r, err, map[string]interface{}{
			"message": getHTTPBodyMessage(w, isUpdated, UPDATED),
			"mazeId":  updatedMazeId,
		})
	}
	return fn
}

func GetMaze() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var maze maze.Maze
		mazeId := stringUtil.ParseUint(chi.URLParam(r, "id"))
		err := mazeService.FindById(&maze, mazeId)
		if err != nil {
			log.Println(err)
		}
		httpUtil.GenerateResponse(w, r, err, maze)
	}
	return fn
}

func getHTTPBodyMessage(w http.ResponseWriter, isSuccess bool, message string) (messageBody string) {
	switch isSuccess {
	case true:
		messageBody = fmt.Sprintf("Successfully %s", message)
		w.WriteHeader(http.StatusOK)
	case false:
		messageBody = fmt.Sprintf("Couldn't %s", message)
		w.WriteHeader(http.StatusBadRequest)
	}
	return
}
