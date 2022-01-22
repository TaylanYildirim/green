package mazeHandler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"green/database"
	"green/handlers"
	"green/models/maze"
	"green/utils/httpUtil"
	"green/utils/stringUtil"
	"io/ioutil"
	"log"
	"net/http"
)

func InsertMaze() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var maze maze.Maze
		err := json.NewDecoder(r.Body).Decode(&maze)
		if err != nil {
			handlers.SetHTTPStatus(w, http.StatusBadRequest, "StatusBadRequest", 0)
			return
		}
		defer r.Body.Close()

		jsonBody, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(jsonBody, &maze)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(maze)

	}
	return fn
}

func DeleteMaze() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {

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
		var maze maze.MongoMaze
		mazeId := stringUtil.ParseUint(chi.URLParam(r, "id"))
		err := database.FindById("maze", bson.M{"mazeId": mazeId}, &maze)
		if err != nil {
			log.Fatal(err)
		}
		httpUtil.JSON(w, r, maze)
	}
	return fn
}
