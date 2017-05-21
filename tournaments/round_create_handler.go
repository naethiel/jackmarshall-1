package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-kit/kit/log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/chibimi/jackmarshall/auth"
	"github.com/julienschmidt/httprouter"
)

func NewCreateRoundHandler(db *mgo.Session, logger log.Logger) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := auth.Context(r)

		// Check if the user is admin or has a valid role
		ok, admin := ctx.User.IsAuthorized(auth.RoleOrga)
		if !ok {
			logger.Log("request_id", ctx.RequestID, "level", "info", "msg", "Invalid roles", "roles", strings.Join(ctx.User.Roles, ", "))
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(fmt.Sprintf("Invalid roles: %s", ctx.User.Roles)))
			return
		}

		collection := db.DB("jackmarshall").C("tournament")
		id := p.ByName("id")
		tournament := Tournament{}

		var err error
		if admin {
			logger.Log("request_id", ctx.RequestID, "level", "debug", "msg", "create round as admin", "tournament_id", id)
			err = collection.FindId(bson.ObjectIdHex(id)).One(&tournament)
		} else {
			err := collection.Find(bson.M{"_id": bson.ObjectIdHex(id), "owner": ctx.User.ID}).One(&tournament)
		}
		if err != nil {
			logger.Log("request_id", ctx.RequestID, "level", "error", "msg", "Unable to create round", "tournament_id", id, "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//create pairings and assign tables
		round := Round{
			Number: len(tournament.Rounds),
			Games:  []Game{},
		}

		if len(tournament.Players) == 0 || len(tournament.Tables) < len(tournament.Players)/2 {
			logger.Log("request_id", ctx.RequestID, "level", "error", "msg", "Unable to create round", "tournament_id", id, "error", "Incorect number of players or tables")
			http.Error(w, "Incorect number of players or tables", http.StatusBadRequest)
			return
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
			logger.Log("request_id", ctx.RequestID, "level", "error", "msg", "Unable to add round", "tournament_id", id, "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tournament)
	}
}
