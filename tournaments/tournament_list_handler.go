package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func NewListTournamentHandler(db *mgo.Session) httprouter.Handle {
	collection := db.DB("jackmarshall").C("tournament")

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		results := []Tournament{}

		err := collection.Find(nil).All(&results)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(results)
	}
}
