package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/HouzuoGuo/tiedot/data"
	"github.com/julienschmidt/httprouter"
)

func NewGetResultsHandler(database *data.Collection) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		id, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tournament := Tournament{}

		doc := database.Read(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(doc, &tournament)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		result := tournament.getResults()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}
