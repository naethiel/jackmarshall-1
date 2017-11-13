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
	. "github.com/chibimi/jackmarshall/tournaments"
	"github.com/julienschmidt/httprouter"
)

// NewListTournamentHandler return all tournament matching the user id.
func NewListTournamentHandler(db *mgo.Session, logger log.Logger) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
		results := []Tournament{}

		var err error
		if admin {
			logger.Log("request_id", ctx.RequestID, "level", "debug", "msg", "list tournament as admin")
			err = collection.Find(nil).All(&results)
		} else {
			err = collection.Find(bson.M{"owner": ctx.User.ID}).All(&results)
		}
		if err != nil {
			logger.Log("request_id", ctx.RequestID, "level", "error", "msg", "Unable to list tournaments", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(results)
	}
}
