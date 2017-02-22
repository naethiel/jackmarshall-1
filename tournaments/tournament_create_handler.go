package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
)

func NewCreateTournamentHandler(db *mgo.Session) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		collection := db.DB("jackmarshall").C("tournament")
		tournament := NewTournament()
		err := json.NewDecoder(r.Body).Decode(&tournament)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userID, err := strconv.ParseInt(p.ByName("userId"), 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tournament.Owner = userID

		tournament.ID = bson.NewObjectId()
		err = collection.Insert(&tournament)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tournament.ID)
	}
}
