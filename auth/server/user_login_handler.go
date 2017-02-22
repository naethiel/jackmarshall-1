package main

import (
	"encoding/json"
	"net/http"

	"github.com/chibimi/jackmarshall/auth"
	"github.com/elwinar/token"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"menteslibres.net/gosexy/redis"
)

func NewUserLoginHandler(db *redis.Client, c Configuration) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		// Decode the credentials from the payload
		credentials := struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		}{}
		err := decoder.Decode(&credentials)
		if err != nil {
			http.Error(w, "malformed request payload: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Get the corresponding user ID
		ID, err := db.HGet("users", credentials.Login)
		switch err {
		case nil:
			break
		case redis.ErrNilReply:
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		default:
			http.Error(w, "unable to get user id: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Get the user
		raw, err := db.Get("user:" + ID)
		if err != nil {
			http.Error(w, "unable to get user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Unmarshall the user
		user := auth.User{}
		err = json.Unmarshal([]byte(raw), &user)
		if err != nil {
			http.Error(w, "unable to unmarshal user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Check the password, mind the secure comparison
		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)) != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// Initialize the claims to be used
		claims := token.Claims{
			"user": user,
		}

		// Generate the token
		token, err := token.SignHS256(claims, []byte(c.Secret))
		if err != nil {
			http.Error(w, "unable to generate the token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Send the token
		w.Write([]byte(token))
	}
}
