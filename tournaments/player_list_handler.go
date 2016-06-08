package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func NewListPlayerHandler(db *mgo.Session) httprouter.Handle {
	collection := db.DB("jackmarshall").C("tournament")

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		var results []string

		if p.ByName("root") == "ok" {
			err := collection.Find(nil).Distinct("players.name", &results)
			if err != nil {
				panic(err)
			}
		} else {
			userID, _ := strconv.ParseInt(p.ByName("userId"), 10, 64)
			err := collection.Find(bson.M{"owner": userID}).Distinct("players.name", &results)
			if err != nil {
				panic(err)
			}
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(results)
	}
}
