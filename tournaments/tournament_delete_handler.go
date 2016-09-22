package main

import (
	"net/http"
	"strconv"

	"github.com/HouzuoGuo/tiedot/db"
	"github.com/julienschmidt/httprouter"
)

func NewDeleteTournamentHandler(database *db.DB) httprouter.Handle {
	collection := database.Use("Tournaments")

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		id, err := strconv.Atoi(p.ByName("id"))

		err = collection.Delete(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
