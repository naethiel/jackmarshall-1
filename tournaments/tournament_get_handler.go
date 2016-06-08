package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func NewGetTournamentHandler(db *mgo.Session) httprouter.Handle {
	collection := db.DB("jackmarshall").C("tournament")

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		id := p.ByName("id")

		result := Tournament{}

		err := collection.FindId(bson.ObjectIdHex(id)).One(&result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}
