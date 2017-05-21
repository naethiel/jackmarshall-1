package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/chibimi/jackmarshall/auth"
	"github.com/elwinar/token"
	"github.com/go-kit/kit/log"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"menteslibres.net/gosexy/redis"
)

func NewUserUpdateHandler(db *redis.Client, logger log.Logger, c Configuration) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Check that the token is valid
		claims, err := token.ParseHS256(r.Header.Get("Authorization"), []byte(c.Secret))
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// Get the id from the parameters
		id, err := strconv.ParseInt(p.ByName("id"), 10, 64)
		if err != nil {
			http.Error(w, "invalid id: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Get the user from the database
		user, err := auth.NewUserFromDatabase(db, id)
		if err != nil {
			http.Error(w, "unable to get user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Get the requester from the claims
		requester, err := auth.NewUserFromClaims(claims["user"])
		if err != nil {
			http.Error(w, "invalid claim:"+err.Error(), http.StatusUnauthorized)
			return
		}

		// If the requester isn't root, it must be the same user than the updated one
		if requester.HasRole("root") && requester.ID != user.ID {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// Decode the request
		request := struct {
			Password *string  `json:"password"`
			Roles    []string `json:"roles"`
		}{}
		err = json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "malformed request payload: "+err.Error(), http.StatusBadRequest)
			return
		}

		// If the request include a new set of role and the requester is root
		if requester.HasRole("root") && request.Roles != nil {
			user.Roles = request.Roles
		}

		// If the request include a new password
		if request.Password != nil {
			// Hash the password
			password, err := bcrypt.GenerateFromPassword([]byte(*request.Password), c.PasswordCost)
			if err != nil {
				http.Error(w, "unable to hash the new password for the user: "+err.Error(), http.StatusInternalServerError)
				return
			}
			user.Password = string(password)
		}
		user.Login = strings.ToLower(user.Login)

		// Insert the json-encoded user in the database
		raw, _ := json.Marshal(user)
		_, err = db.Set(fmt.Sprintf("user:%d", user.ID), string(raw))
		if err != nil {
			http.Error(w, "unable to update the value for the user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Report success
		w.Write([]byte("OK"))
	}
}
