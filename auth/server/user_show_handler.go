package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/chibimi/jackmarshall/auth"
	"github.com/go-kit/kit/log"
	"github.com/julienschmidt/httprouter"
	"menteslibres.net/gosexy/redis"
)

func NewUserShowHandler(db *redis.Client, logger log.Logger) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id, err := strconv.ParseInt(p.ByName("id"), 10, 64)
		if err != nil {
			logger.Log("level", "error", "msg", "unable to parse id", "error", err, "id", p.ByName("id"))
			http.Error(w, "invalid id: "+err.Error(), http.StatusBadRequest)
			return
		}

		user, err := auth.NewUserFromDatabase(db, id)
		if err != nil {
			logger.Log("level", "error", "msg", "unable to get user from db", "error", err, "id", id)
			http.Error(w, "unable to get user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Ensure the password hash isn't disclosed
		user.Password = ""
		user.Secret = ""

		// Send the user
		response, _ := json.Marshal(user) // Skipping the error because the user was just unmarshalled
		w.Write([]byte(response))
	}
}
