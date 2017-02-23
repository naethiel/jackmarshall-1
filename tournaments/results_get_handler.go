package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
)

func NewGetResultsHandler(db *mgo.Session) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		collection := db.DB("jackmarshall").C("tournament")

		id := p.ByName("id")

		tournament := Tournament{}

		if p.ByName("root") == "ok" {
			err := collection.FindId(bson.ObjectIdHex(id)).One(&tournament)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			userID, _ := strconv.ParseInt(p.ByName("userId"), 10, 64)
			err := collection.Find(bson.M{"_id": bson.ObjectIdHex(id), "owner": userID}).One(&tournament)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		result := tournament.getResults()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}
