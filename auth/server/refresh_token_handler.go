package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/chibimi/jackmarshall/auth"
	"github.com/elwinar/token"
	"github.com/julienschmidt/httprouter"
	"menteslibres.net/gosexy/redis"
)

func NewRefreshTokenHandler(db *redis.Client, c Configuration) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "unable to read token: "+err.Error(), http.StatusBadRequest)
			return
		}

		chunks := strings.Split(string(data), ".")
		if len(chunks) != 3 {
			http.Error(w, "malformed token: "+err.Error(), http.StatusBadRequest)
			return
		}

		var claims token.Claims
		decodedClaims, err := base64.URLEncoding.DecodeString(chunks[1])
		if err != nil {
			http.Error(w, "unable to decode header: "+err.Error(), http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(decodedClaims, &claims)
		if err != nil {
			http.Error(w, "unable to unmarshall header: "+err.Error(), http.StatusBadRequest)
			return
		}

		userID, found := claims["userID"]
		if !found {
			http.Error(w, "invalid token: userID not found.", http.StatusBadRequest)
			return
		}
		id, ok := userID.(float64)
		if !ok {
			http.Error(w, "invalid token: userID must be a Numerical.", http.StatusBadRequest)
			return
		}
		user, err := auth.NewUserFromDatabase(db, int64(id))
		if err != nil {
			http.Error(w, "unable to get user: "+err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := token.ParseHS256(string(data), []byte(user.Secret)); err != nil {
			http.Error(w, "invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}
		// Ensure the password hash isn't disclosed
		user.Password = ""
		user.Secret = ""

		// Initialize the claims to be used
		exp := time.Now().Add(1 * time.Minute).Unix()
		claims = token.Claims{
			"user": user,
			"exp":  exp,
		}

		// Generate the new token
		token, err := token.SignHS256(claims, []byte(c.Secret))
		if err != nil {
			http.Error(w, "unable to generate the token: "+err.Error(), http.StatusInternalServerError)
			return
		}
		auth := Auth{
			Token:      token,
			Expiration: exp,
		}

		res, err := json.Marshal(auth)
		if err != nil {
			http.Error(w, "unable to generate the auth: "+err.Error(), http.StatusInternalServerError)
			return
		}
		// Send the tokens
		w.Write(res)
	}
}

func NewInvalidateRefreshTokenHandler(db *redis.Client, c Configuration) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "unable to read token: "+err.Error(), http.StatusBadRequest)
			return
		}

		chunks := strings.Split(string(data), ".")
		if len(chunks) != 3 {
			http.Error(w, "malformed token: "+err.Error(), http.StatusBadRequest)
			return
		}

		var claims token.Claims
		decodedClaims, err := base64.URLEncoding.DecodeString(chunks[1])
		if err != nil {
			http.Error(w, "unable to decode header: "+err.Error(), http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(decodedClaims, &claims)
		if err != nil {
			http.Error(w, "unable to unmarshall header: "+err.Error(), http.StatusBadRequest)
			return
		}

		userID, found := claims["userID"]
		if !found {
			http.Error(w, "invalid token: userID not found.", http.StatusBadRequest)
			return
		}
		id, ok := userID.(float64)
		if !ok {
			http.Error(w, "invalid token: userID must be a Numerical.", http.StatusBadRequest)
			return
		}
		user, err := auth.NewUserFromDatabase(db, int64(id))
		if err != nil {
			http.Error(w, "unable to get user: "+err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := token.ParseHS256(string(data), []byte(c.Secret)); err != nil {
			http.Error(w, "invalid token: "+err.Error(), http.StatusBadRequest)
			return
		}

		user.Secret = RandomString(8)

		raw, _ := json.Marshal(user)
		_, err = db.Set(fmt.Sprintf("user:%d", user.ID), string(raw))
		if err != nil {
			http.Error(w, "unable to update user secret: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Report success
		w.Write([]byte("OK"))
	}
}
