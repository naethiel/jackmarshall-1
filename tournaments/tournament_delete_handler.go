package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/chibimi/jackmarshall/auth"
	"github.com/go-kit/kit/log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
)

func NewDeleteTournamentHandler(db *mgo.Session, logger log.Logger) httprouter.Handle {
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

		var err error
		if admin {
			logger.Log("request_id", ctx.RequestID, "level", "debug", "msg", "delete tournament as admin", "tournament_id", id)
			err = collection.RemoveId(bson.ObjectIdHex(id))
		} else {
			err = collection.Remove(bson.M{"_id": bson.ObjectIdHex(id), "owner": ctx.User.ID})
		}
		if err != nil {
			logger.Log("request_id", ctx.RequestID, "level", "error", "msg", "Unable to delete tournament", "tournament_id", id, "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
