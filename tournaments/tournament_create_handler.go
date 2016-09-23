package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/HouzuoGuo/tiedot/data"
	"github.com/julienschmidt/httprouter"
)

func NewCreateTournamentHandler(database *data.Collection) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		tournament := NewTournament()
		err := json.NewDecoder(r.Body).Decode(&tournament)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data, err := json.Marshal(tournament)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		id, err := database.Insert(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(strconv.Itoa(id))
	}
}
