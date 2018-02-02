package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"

	"github.com/chibimi/jackmarshall/tournaments"
)

func displayRound(r tournaments.Round) {
	fmt.Println("Round", r.Number)
	for _, g := range r.Games {
		fmt.Printf("Table %s: %s (%d) vs %s (%d)\n", g.TableID, g.Results[0].PlayerID, g.Results[0].VictoryPoints, g.Results[1].PlayerID, g.Results[1].VictoryPoints)
	}
}

func main() {
	pairing, origin, bye, ssapp := 0, 0, 0, 0
	ErrPairing := []Error{}
	for i := 0; i < 100; i++ {
		t := tournaments.NewTestTournament(60, 32, 6, false)
		nbRound := 6
		for j := 0; j < nbRound; j++ {
			round := t.GetNextRound()
			round.Number = j
			for k, _ := range round.Games {
				if round.Games[k].Results[0].Bye {
					continue
				}
				round.Games[k].Results[rand.Intn(2)].VictoryPoints = 1
			}
			t.Rounds = append(t.Rounds, round)
			p, o, b, s := t.CheckRound(round.Number)
			if len(p) != 0 {
				pairing += len(p)
				ErrPairing = append(ErrPairing, Error{
					Type:       "pairing",
					Error:      p,
					Round:      round.Number,
					Tournament: *t,
				})
				fmt.Println("==> Pairing at round", j, p)
			}
			if len(o) != 0 {
				origin += len(o)
				fmt.Println("==> Origin", j, o)
			}
			if len(b) != 0 {
				bye += len(b)
				fmt.Println("==> Bye", j, b)
			}
			if len(s) != 0 {
				ssapp += len(s)
				fmt.Println("==> SousApp", j, s)
			}
		}
	}
	data, _ := json.Marshal(ErrPairing)
	ioutil.WriteFile("errPair.json", data, 0664)
	fmt.Println(pairing, origin, bye, ssapp)
}

type Error struct {
	Type       string
	Error      []string
	Round      int
	Tournament tournaments.Tournament
}
