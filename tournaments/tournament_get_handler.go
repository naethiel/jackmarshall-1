package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/HouzuoGuo/tiedot/db"
	"github.com/julienschmidt/httprouter"
)

func NewGetTournamentHandler(database *db.DB) httprouter.Handle {
	collection := database.Use("Tournaments")

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		id, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result := Tournament{}

		doc, err := collection.Read(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result.Id = p.ByName("id")
		result.Name = doc["name"].(string)
		result.Date, err = time.Parse(time.RFC3339, doc["date"].(string))
		result.Format = int(doc["format"].(float64))
		result.Slots = int(doc["slots"].(float64))
		result.FeeAmount = doc["fee_amount"].(float64)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}
