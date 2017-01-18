package main

import (
	"encoding/json"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
)

func NewCreateRoundHandler(db *mgo.Session) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		collection := db.DB("jackmarshall").C("tournament")

		id := p.ByName("id")

		tournament := Tournament{}

		err := collection.FindId(bson.ObjectIdHex(id)).One(&tournament)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//create pairings and assign tables
		round := Round{
			Number: len(tournament.Rounds),
			Games:  []Game{},
		}

		var pairings = CreatePairs(tournament.Players, tournament, &round)
		createRound(pairings, tournament.Tables, &round)

		for i, _ := range round.Games {
			round.Games[i].Results[0].Player.Games = nil
			round.Games[i].Results[1].Player.Games = nil
		}

		tournament.Rounds = append(tournament.Rounds, round)

		err = collection.UpdateId(bson.ObjectIdHex(id), &tournament)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tournament)
	}
}
