package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func NewCreateRoundHandler(db *mgo.Session) httprouter.Handle {
	collection := db.DB("jackmarshall").C("tournament")
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		//Get tournament
		id := p.ByName("id")
		tournament := Tournament{}
		err := collection.FindId(bson.ObjectIdHex(id)).One(&tournament)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, player := range tournament.Players {
			for r, round := range tournament.Rounds {
				for g, game := range round.Games {
					for _, result := range game.Results {
						if result.Player == player.Name {
							player.Games = append(player.Games, &tournament.Rounds[r].Games[g])
						}
					}
				}

			}
		}

		//create pairings and assign tables
		round := Round{
			Games: []Game{},
		}
		//	fmt.Println(players)
		var pairings = CreatePairs(tournament.Players, tournament, &round)

		CreateRound(pairings, tournament.Tables, &round)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(round)
	}
}
