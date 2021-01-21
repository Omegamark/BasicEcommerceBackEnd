package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type MongoDataStore interface {
	GetStuff() ([]map[string]interface{}, error)
	InsertStuff(thing map[string]interface{}) error
}

type MongoService struct {
	Version string
	DBCon   MongoDataStore
}

func (m *MongoService) Routes(r *mux.Router) *mux.Router {
	r.HandleFunc("/getGoods", m.GetGoods).Methods("GET")
	r.HandleFunc("/postGoods", m.InsertGood).Methods("POST")
	return r
}

func (m *MongoService) InsertGood(w http.ResponseWriter, r *http.Request) {
	thing := map[string]interface{}{}

	err := json.NewDecoder(r.Body).Decode(&thing)
	if err != nil {
		fmt.Println(err)
	}

	m.DBCon.InsertStuff(thing)
}

func (m *MongoService) GetGoods(w http.ResponseWriter, r *http.Request) {
	result, err := m.DBCon.GetStuff()

	responsePayload, err := json.Marshal(result)
	if err != nil {
		fmt.Errorf("Handler failed with error: %s", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responsePayload)
}
