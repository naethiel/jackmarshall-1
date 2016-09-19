package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func NewCreateTournamentHandler(db *mgo.Session) httprouter.Handle {
	collection := db.DB("jackmarshall").C("tournament")

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		var tournament Tournament
		err := json.NewDecoder(r.Body).Decode(&tournament)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = collection.Insert(&tournament)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
