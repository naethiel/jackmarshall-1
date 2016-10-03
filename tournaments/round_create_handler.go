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

		// //get every player's games
		// for _, player := range tournament.Players {
		// 	for r, round := range tournament.Rounds {
		// 		for g, game := range round.Games {
		// 			for _, result := range game.Results {
		// 				if result.Player == player.Name {
		// 					player.Games = append(player.Games, &tournament.Rounds[r].Games[g])
		// 				}
		// 			}
		// 		}
		// 	}
		// }

		//create pairings and assign tables
		round := Round{
			Games: []Game{},
		}
		var pairings = CreatePairs(tournament.Players, tournament, &round)

		CreateRound(pairings, tournament.Tables, &round)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(round)
	}
}
