package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func NewDeleteScenarioHandler(db *mgo.Session) httprouter.Handle {
	collection := db.DB("jackmarshall").C("scenario")

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		id := p.ByName("id")

		err := collection.RemoveId(bson.ObjectIdHex(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
