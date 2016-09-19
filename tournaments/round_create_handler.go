package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"time"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func NewCreateRoundHandler(db *mgo.Session) httprouter.Handle {
	collection := db.DB("jackmarshall").C("tournament")
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		//Get tournament
		id := p.ByName("id")
		tournament := Tournament{}
		err := collection.FindId(bson.ObjectIdHex(id)).One(&tournament)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//Get tables
		tables := []string{}
		err = collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).Distinct("tables.name", &tables)
		if err != nil {
			panic(err)
		}

		//Get active players
		players := []string{}
		err = collection.Find(bson.M{"_id": bson.ObjectIdHex(id), "players.leave": false}).Distinct("players.name", &players)
		if err != nil {
			panic(err)
		}

		//shuffle palyers list
		shuffle(players)

		//Sort player list by victory points
		sortPlayersByVictoryPoints(players, tournament)

		fmt.Println("players : ", players)
		fmt.Println("tables : ", tables)

		//create pairings
		pairings := CreatePairings(players, tournament)

		fmt.Println("pairings : ", pairings)

		//assign tables
		min := 50
		minScore := &min
		solution := assignTables(make([]string, len(tables)), 0, tables, len(tables), pairings, tournament, make([]string, len(tables)), minScore)

		fmt.Println("solution : ", solution)

		round := NewRoundFromPairing(pairings, solution)

		fmt.Println("round : ", round)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(round)
	}
}

func shuffle(slc []string) {
	rand.Seed(time.Now().UnixNano())
	N := len(slc)
	for i := 0; i < N; i++ {
		// choose index uniformly in [i, N-1]
		r := i + rand.Intn(N-i)
		slc[r], slc[i] = slc[i], slc[r]
	}
}

func sortPlayersByVictoryPoints(players []string, tournament Tournament) {
	score := PlayerScoreList{}
	for _, p := range players {
		score = append(score, PlayerScore{p, tournament.getVictoryPoints(p)})
	}

	sort.Sort(sort.Reverse(score))

	players = players[:0]
	for _, v := range score {
		players = append(players, v.Player)
	}
}

func assignTables(attempt []string, position int, remainingTables []string, nbTables int, pairings []Pairing, tournament Tournament, solution []string, minScore *int) []string {
	if len(remainingTables) == 0 {
		*minScore = calculateScore(attempt, pairings, tournament)
		copy(solution, attempt)
		if *minScore == 0 {
			return solution
		}
	}

	for i, v := range remainingTables {
		attempt[position] = v
		if calculateScore(attempt, pairings, tournament) >= *minScore {
			attempt = append(attempt[:position], make([]string, nbTables-position)...)
		} else {
			paramRemainingTables := make([]string, len(remainingTables))
			copy(paramRemainingTables, remainingTables)

			paramRemainingTables = append(paramRemainingTables[:i], paramRemainingTables[i+1:]...)

			assignTables(attempt, position+1, paramRemainingTables, nbTables, pairings, tournament, solution, minScore)
		}
	}
	return solution
}

func calculateScore(attempt []string, pairings []Pairing, tournament Tournament) (res int) {

	res = 0
	for i, v := range attempt {
		player1Tables := tournament.getTablesPlayed(pairings[i].Player1)
		player2Tables := tournament.getTablesPlayed(pairings[i].Player2)

		if index := contains(player1Tables, v); index != -1 {
			res += 10
			if index == (len(player1Tables) - 1) {
				res += 5
			}
		}
		if index := contains(player2Tables, v); index != -1 {
			res += 10
			if index == (len(player2Tables) - 1) {
				res += 5
			}
		}
	}
	return
}

func contains(s []string, e string) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}
