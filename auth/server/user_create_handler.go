package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/chibimi/jackmarshall/auth"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"menteslibres.net/gosexy/redis"
)

func NewUserCreateHandler(db *redis.Client, logger log.Logger, c Configuration) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		// Decode the use from the payload
		user := auth.User{}
		err := decoder.Decode(&user)
		if err != nil {
			logger.Log("level", "error", "msg", "unable to decode body", "error", err)
			http.Error(w, "malformed request payload: "+err.Error(), http.StatusBadRequest)
			return
		}
		user.Login = strings.ToLower(user.Login)

		// Ensure the roles are always empty when creating a new user, so that
		// they can only be set by administrators in a protected endpoint.
		user.Roles = nil

		// Check if the user already exists
		_, err = db.HGet("users", user.Login)
		if err == nil {
			logger.Log("level", "error", "msg", "unable to create user", "error", "user already exists", "username", user.Login)
			http.Error(w, "login "+user.Login+" already exists", http.StatusBadRequest)
			return
		}

		// Hash the password
		password, err := bcrypt.GenerateFromPassword([]byte(user.Password), c.PasswordCost)
		if err != nil {
			logger.Log("level", "error", "msg", "unable to hash password", "error", err)
			http.Error(w, "unable to hash the password for the new user: "+err.Error(), http.StatusInternalServerError)
			return
		}
		user.Password = string(password)

		// Generate random Secret
		user.Secret = RandomString(8)

		// Get an ID for the new user
		ID, err := db.Incr("users_max")
		if err != nil {
			logger.Log("level", "error", "msg", "unable to generate id", "error", err)
			http.Error(w, "unable to generate an id for the new user: "+err.Error(), http.StatusInternalServerError)
			return
		}
		user.ID = ID

		// Insert the new user in the users index
		_, err = db.HSet("users", user.Login, user.ID)
		if err != nil {
			logger.Log("level", "error", "msg", "unable to index user", "error", err)
			http.Error(w, "unable to index the new user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Insert the json-encoded user in the database
		raw, _ := json.Marshal(user)
		_, err = db.Set(fmt.Sprintf("user:%d", user.ID), string(raw))
		if err != nil {
			logger.Log("level", "error", "msg", "unable to save user in db", "error", err)
			http.Error(w, "unable to save the new user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Report success
		w.Write([]byte("OK"))
	}
}

func NewOrganizerCreateHandler(db *redis.Client, logger log.Logger, c Configuration) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		// Decode the use from the payload
		user := auth.User{}
		err := decoder.Decode(&user)
		if err != nil {
			logger.Log("level", "error", "msg", "unable to decode body", "error", err)
			http.Error(w, "malformed request payload: "+err.Error(), http.StatusBadRequest)
			return
		}
		user.Login = strings.ToLower(user.Login)

		// Ensure the roles are always empty when creating a new user, so that
		// they can only be set by administrators in a protected endpoint.
		user.Roles = []string{auth.RoleOrga}

		// Check if the user already exists
		_, err = db.HGet("users", user.Login)
		if err == nil {
			logger.Log("level", "error", "msg", "unable to create user", "error", "user already exists", "username", user.Login)
			http.Error(w, "login "+user.Login+" already exists", http.StatusBadRequest)
			return
		}

		// Hash the password
		password, err := bcrypt.GenerateFromPassword([]byte(user.Password), c.PasswordCost)
		if err != nil {
			logger.Log("level", "error", "msg", "unable to hash password", "error", err)
			http.Error(w, "unable to hash the password for the new user: "+err.Error(), http.StatusInternalServerError)
			return
		}
		user.Password = string(password)

		// Generate random Secret
		user.Secret = RandomString(8)

		// Get an ID for the new user
		ID, err := db.Incr("users_max")
		if err != nil {
			logger.Log("level", "error", "msg", "unable to generate id", "error", err)
			http.Error(w, "unable to generate an id for the new user: "+err.Error(), http.StatusInternalServerError)
			return
		}
		user.ID = ID

		// Insert the new user in the users index
		_, err = db.HSet("users", user.Login, user.ID)
		if err != nil {
			logger.Log("level", "error", "msg", "unable to index user", "error", err)
			http.Error(w, "unable to index the new user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Insert the json-encoded user in the database
		raw, _ := json.Marshal(user)
		_, err = db.Set(fmt.Sprintf("user:%d", user.ID), string(raw))
		if err != nil {
			logger.Log("level", "error", "msg", "unable to save user in db", "error", err)
			http.Error(w, "unable to save the new user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Report success
		w.Write([]byte("OK"))
	}
}

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

func RandomString(len int) string {
	b := make([]byte, len)
	r.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
