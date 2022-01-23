package mazeHandler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"green/database"
	"green/handlers/shortestPath"
	"green/models/maze"
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

		isInserted, err := database.InsertOne("maze", bson.M{"maze": newMaze.Maze, "mazeId": newMaze.MazeId})

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
		isDeleted, err := database.DeleteOne("maze", bson.M{"mazeId": mazeId})
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

		isUpdated, err := database.UpdateOne("maze", bson.M{"mazeId": updatedMazeId}, bson.M{"maze": updatedMaze.Maze})

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
		shortestPath.GetShortestPath()
		mazeId := stringUtil.ParseUint(chi.URLParam(r, "id"))
		err := database.FindById("maze", bson.M{"mazeId": mazeId}, &maze)
		if err != nil {
			log.Fatal(err)
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
