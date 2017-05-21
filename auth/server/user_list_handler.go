package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/log"
	"github.com/julienschmidt/httprouter"
	"menteslibres.net/gosexy/redis"
)

func NewUserListHandler(db *redis.Client, logger log.Logger) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		hash, err := db.HGetAll("users")
		if err != nil {
			logger.Log("level", "error", "msg", "unable to retrieve user indexes", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("unable to retrieve users indexes"))
			return
		}

		index := make(map[string]int64)
		for i := 0; i < len(hash); i += 2 {
			ID, err := strconv.ParseInt(hash[i+1], 10, 64)
			if err != nil {
				logger.Log("level", "error", "msg", "unable to parse hash", "error", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("unable to parse hash"))
				return
			}
			index[hash[i]] = ID
		}

		response, _ := json.Marshal(index)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}
}
