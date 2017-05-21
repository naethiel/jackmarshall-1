package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/chibimi/jackmarshall/auth"
	"github.com/go-kit/kit/log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
)

func NewCreateTournamentHandler(db *mgo.Session, logger log.Logger) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := auth.Context(r)

		// Check if the user is admin or has a valid role
		ok, _ := ctx.User.IsAuthorized(auth.RoleOrga)
		if !ok {
			logger.Log("request_id", ctx.RequestID, "level", "info", "msg", "Invalid roles", "roles", strings.Join(ctx.User.Roles, ", "))
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(fmt.Sprintf("Invalid roles: %s", ctx.User.Roles)))
			return
		}

		collection := db.DB("jackmarshall").C("tournament")
		tournament := NewTournament()
		err := json.NewDecoder(r.Body).Decode(&tournament)
		if err != nil {
			logger.Log("request_id", ctx.RequestID, "level", "error", "msg", "Unable to decode body", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tournament.Owner = ctx.User.ID
		tournament.ID = bson.NewObjectId()

		err = collection.Insert(&tournament)
		if err != nil {
			logger.Log("request_id", ctx.RequestID, "level", "error", "msg", "Unable to insert tournament in database", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tournament.ID)
	}
}
