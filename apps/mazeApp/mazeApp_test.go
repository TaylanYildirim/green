package mazeApp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"green/config"
	"green/database"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	config.FileName = "../../conf.yaml"
	database.InitDB()
}

func TestDeleteMaze(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/maze/%s", "0"), nil)
	w := httptest.NewRecorder()

	DeleteMaze()(w, req)

	res := w.Result()
	defer res.Body.Close()
	if res.Status != "200 OK" {
		t.Errorf("err: response status %v", res.Status)
	}
}

func TestGetMaze(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/maze/%s", "1"), nil)
	w := httptest.NewRecorder()

	GetMaze()(w, req)

	res := w.Result()
	defer res.Body.Close()
	if res.Status == "404 NotFound" {
		t.Errorf("err: response status %v", res.Status)
	}
}
func TestInsertMaze(t *testing.T) {
	reqBody := map[string][][]int32{
		"maze": {
			{1, 1, 1, 1, 0, 1, 1, 1},
			{1, 0, 0, 0, 0, 0, 0, 1},
			{1, 0, 1, 1, 1, 0, 1, 1},
			{1, 0, 0, 0, 1, 0, 0, 1},
			{1, 1, 1, 0, 1, 1, 0, 1},
			{1, 0, 0, 0, 1, 0, 0, 1},
			{1, 0, 1, 1, 1, 0, 1, 1},
			{1, 0, 0, 2, 0, 0, 0, 1},
			{1, 1, 1, 1, 1, 1, 1, 1},
		},
	}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/maze", bytes.NewReader(body))
	w := httptest.NewRecorder()

	InsertMaze()(w, req)

	res := w.Result()
	defer res.Body.Close()
	if res.Status != "200 OK" {
		t.Errorf("err: response status %v", res.Status)
	}
}
func TestUpdateMaze(t *testing.T) {
	reqBody := map[string][][]int32{
		"maze": {
			{1, 1, 1, 1, 0, 1, 1, 1},
			{1, 0, 0, 0, 0, 0, 0, 1},
			{1, 0, 1, 1, 1, 0, 1, 1},
			{1, 0, 0, 0, 1, 0, 0, 1},
			{1, 1, 1, 0, 1, 1, 0, 1},
			{1, 1, 0, 0, 1, 0, 0, 1},
			{0, 0, 1, 1, 1, 0, 1, 1},
			{1, 0, 0, 2, 0, 0, 0, 1},
			{1, 1, 1, 1, 1, 1, 1, 1},
		},
	}
	updatedBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/maze/%s", "0"), bytes.NewReader(updatedBody))
	w := httptest.NewRecorder()

	InsertMaze()(w, req)

	res := w.Result()
	defer res.Body.Close()
	if res.Status != "200 OK" {
		t.Errorf("err: response status %v", res.Status)
	}

}
