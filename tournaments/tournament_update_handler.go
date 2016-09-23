package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/HouzuoGuo/tiedot/data"
	"github.com/julienschmidt/httprouter"
)

func NewUpdateTournamentHandler(database *data.Collection) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		id, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var tournament Tournament
		err = json.NewDecoder(r.Body).Decode(&tournament)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data, err := json.Marshal(tournament)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		id, err = database.Update(id, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(strconv.Itoa(id))
	}
}
