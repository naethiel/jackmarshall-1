package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/HouzuoGuo/tiedot/data"
	"github.com/julienschmidt/httprouter"
)

func NewListTournamentHandler(database *data.Collection) httprouter.Handle {
	// collection := database.Use("Tournaments")

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		results := []Tournament{}

		database.ForEachDoc(func(id int, doc []byte) (willMoveOn bool) {
			var tournament Tournament
			err := json.Unmarshal(doc, &tournament)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			tournament.Id = strconv.Itoa(id)
			results = append(results, tournament)
			return true // move on to the next document OR
		})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(results)
	}
}
