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

func NewGetResultsHandler(db *mgo.Session, logger log.Logger) httprouter.Handle {
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
			logger.Log("request_id", ctx.RequestID, "level", "debug", "msg", "get results as admin", "tournament_id", id)
			err = collection.FindId(bson.ObjectIdHex(id)).One(&tournament)
		} else {
			err = collection.Find(bson.M{"_id": bson.ObjectIdHex(id), "owner": ctx.User.ID}).One(&tournament)
		}
		if err != nil {
			logger.Log("request_id", ctx.RequestID, "level", "error", "msg", "Unable to get results", "tournament_id", id, "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := tournament.getResults()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}
