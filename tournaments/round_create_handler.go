package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/HouzuoGuo/tiedot/data"
	"github.com/julienschmidt/httprouter"
)

func NewCreateRoundHandler(database *data.Collection) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		id, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tournament := Tournament{}

		doc := database.Read(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(doc, &tournament)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		tournament.Id = p.ByName("id")

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

		data, err := json.Marshal(tournament)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = database.Update(id, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tournament)
	}
}
