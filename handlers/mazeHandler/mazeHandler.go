package mazeHandler

import (
	"encoding/json"
	"green/handlers"
	"net/http"
)

type Maze struct {
	Maze [][]uint8 `json:"maze"`
}

func InsertMaze() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var jsonMaze Maze
		err := json.NewDecoder(r.Body).Decode(&jsonMaze)
		if err != nil {
			handlers.SetHTTPStatus(w, http.StatusBadRequest, "StatusBadRequest", 0)
			return
		}

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

	}
	return fn
}
