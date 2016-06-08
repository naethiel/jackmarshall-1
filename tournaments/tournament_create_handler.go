package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/validator.v2"
)

//TODO : verifier qu'on ne peut pas creer / update pour un autre owner en accedant directment depuis l'api
func NewCreateTournamentHandler(db *mgo.Session) httprouter.Handle {
	collection := db.DB("jackmarshall").C("tournament")

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		var tournament Tournament
		err := json.NewDecoder(r.Body).Decode(&tournament)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		userID, _ := strconv.Atoi(p.ByName("userId"))
		tournament.Owner = userID

		validator.SetTag("create")
		if err = validator.Validate(tournament); err != nil {
			http.Error(w, "invalid request payload "+err.Error(), http.StatusBadRequest)
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
