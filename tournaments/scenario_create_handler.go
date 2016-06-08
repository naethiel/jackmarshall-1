package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/validator.v2"
)

func NewCreateScenarioHandler(db *mgo.Session) httprouter.Handle {
	collection := db.DB("jackmarshall").C("scenario")

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		var scenario Scenario
		err := json.NewDecoder(r.Body).Decode(&scenario)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err = validator.Validate(scenario); err != nil {
			http.Error(w, "invalid request payload "+err.Error(), http.StatusBadRequest)
			return
		}

		err = collection.Insert(&scenario)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
