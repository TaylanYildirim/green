package mazeHandler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"green/database"
	"green/models/maze"
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

		isInserted, err := database.InsertOne("maze", bson.M{"maze": newMaze.Maze, "mazeId": newMaze.MazeId})

		message := "successfully inserted"
		if !isInserted {
			message = "couldn't inserted"
			w.WriteHeader(http.StatusBadRequest)
		}
		httpUtil.JSON(w, r, map[string]interface{}{
			"message": message,
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

		message := "successfully deleted"
		if !isDeleted {
			message = "couldn't deleted"
			w.WriteHeader(http.StatusNotFound)
		}

		httpUtil.JSON(w, r, map[string]interface{}{
			"message": message,
			"mazeId":  mazeId,
		})
	}
	return fn
}

func UpdateMaze() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {

	}
	return fn
}

func GetMaze() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var maze maze.Maze
		mazeId := stringUtil.ParseUint(chi.URLParam(r, "id"))
		err := database.FindById("maze", bson.M{"mazeId": mazeId}, &maze)
		if err != nil {
			log.Fatal(err)
		}
		httpUtil.JSON(w, r, maze)
	}
	return fn
}
