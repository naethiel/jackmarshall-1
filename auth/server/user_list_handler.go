package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"menteslibres.net/gosexy/redis"
)

func NewUserListHandler(db *redis.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		hash, err := db.HGetAll("users")
		if err != nil {
			log.Println("unable to retrieve the users index:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("unable to retrieve the users index"))
			return
		}

		index := make(map[string]int64)
		for i := 0; i < len(hash); i += 2 {
			ID, err := strconv.ParseInt(hash[i+1], 10, 64)
			if err != nil {
				log.Println("unable to retrieve the users index:", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("unable to retrieve the users index"))
				return
			}

			index[hash[i]] = ID
		}

		response, _ := json.Marshal(index)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}
}
