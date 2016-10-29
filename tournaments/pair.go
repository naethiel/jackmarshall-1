package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bradfitz/slice"
)

type Pair [2]*Player

func (p Pair) PlayedOn(scenario string) bool {
	if p[0].PlayedOn(scenario) {
		return true
	}

	if p[1].PlayedOn(scenario) {
		return true
	}
	return false
}

func CreatePairs(p []Player, tournament Tournament, r *Round) (pairs []Pair) {
	var players = make([]*Player, 0)

	for i := range p {
		if p[i].Leave != true {
			players = append(players, &p[i])
		}
	}

	for _, player := range players {
		player.Games = make([]*Game, 0)
		for r, round := range tournament.Rounds {
			for g, game := range round.Games {
				for _, result := range game.Results {
					if result.Player.Name == player.Name {
						player.Games = append(player.Games, &tournament.Rounds[r].Games[g])
					}
				}
			}
		}
	}

	timeout := false
	perfectPairing := false
	var fuckers []*Player
	go func() {
		time.Sleep(1 * time.Second)
		timeout = true
	}()

	for !timeout && !perfectPairing {
		//shuffle palyers list
		rand.Seed(time.Now().UnixNano())
		for i := range players {
			var j = rand.Intn(len(players) - 1)
			players[i], players[j] = players[j], players[i]
		}

		//Sort player list by victory points
		slice.Sort(players, func(i, j int) bool {
			return players[i].VictoryPoints() > players[j].VictoryPoints()
		})

		//Odd number of players
		if len(players)%2 != 0 {
			for i := len(players) - 1; i >= 0; i-- {
				if players[i].hadBye() == false || len(tournament.Rounds) == 0 {
					r.Games = append(r.Games, Game{
						Table: Table{
							Name: "",
						},
						Results: [2]Result{
							Result{
								Player:            *players[i],
								VictoryPoints:     1,
								ScenarioPoints:    2,
								DestructionPoints: tournament.Format / 2,
								Bye:               true,
							},
							Result{},
						},
					})
					if i == len(players)-1 {
						players = append(players[0:i])
					} else {
						players = append(players[0:i], players[i+1:]...)
					}
					break
				}
			}
		}
		//players that had already played against available players
		fuckers = fuckers[:0]
	Selection:
		for len(players) > 0 {
			for i, p := range players {
				if players[0].Name == p.Name {
					continue
				}
				if p.PlayedAgainst(players[0].Name) {
					continue
				}
				pairs = append(pairs, Pair{players[0], p})
				players = append(players[1:i], players[i+1:]...)
				continue Selection
			}
			fuckers = append(fuckers, players[0])
			players = players[1:]
		}

		perfectPairing = (len(fuckers) == 0)

	}
	// Create pairs from the fuckers.
	for len(fuckers) != 0 {
		fmt.Printf("FUCKERS : %v\n", fuckers)
		pairs = append(pairs, Pair{fuckers[0], fuckers[1]})
		fuckers = fuckers[2:]
	}

	return
}
