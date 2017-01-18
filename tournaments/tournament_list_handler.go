package main

import (
	"encoding/json"
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"github.com/julienschmidt/httprouter"
)

func NewListTournamentHandler(db *mgo.Session) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		collection := db.DB("jackmarshall").C("tournament")

		results := []Tournament{}

		err := collection.Find(nil).All(&results)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(results)
	}
}
