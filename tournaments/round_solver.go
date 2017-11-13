package tournaments

import (
	"fmt"
	"math/rand"
	"time"
)

func GetFitness(games []Game, debug bool) int {
	fit := 0
	for _, g := range games {
		for _, res := range g.Results {
			if res.Player.PlayedTable(g.Table) {
				if debug {
					fmt.Printf("%s a deja joué sur %s\n", res.Player.Name, g.Table.Name)
				}
				fit += 50
			} else if res.Player.PlayedScenario(g.Table.Scenario) {
				if debug {
					fmt.Printf("%s a deja joué sur %s\n", res.Player.Name, g.Table.Scenario)
				}
				fit += 10
			}
		}
	}
	return fit
}

func generateParent(availableTables []Table, availablePairs []Pair) []Game {
	games := make([]Game, 0, len(availablePairs))
	for len(availablePairs) > 0 {
		index := rand.Intn(len(availableTables))
		games = append(games, Game{
			Table: availableTables[index],
			Results: [2]Result{
				Result{
					Player: *availablePairs[0][0],
				},
				Result{
					Player: *availablePairs[0][1],
				},
			},
		})
		availableTables = append(availableTables[:index], availableTables[index+1:]...)
		availablePairs = availablePairs[1:]
	}

	return games
}

func swap(parent []Game) []Game {
	index1 := rand.Intn(len(parent))
	index2 := rand.Intn(len(parent))

	for ; index1 == index2; index2 = rand.Intn(len(parent)) {
	}
	child := make([]Game, len(parent))
	copy(child, parent)
	child[index1].Results = parent[index2].Results
	child[index2].Results = parent[index1].Results

	return child
}

func GetBest(availableTables []Table, availablePairs []Pair, delay time.Duration) []Game {
	bestParent := generateParent(availableTables, availablePairs)
	score := GetFitness(bestParent, false)
	bestScore := score

	timeout := false
	go func() {
		time.Sleep(delay * time.Second)
		timeout = true
	}()

	for !timeout && bestScore != 0 {
		child := swap(bestParent)
		score = GetFitness(child, false)
		if score < bestScore {
			bestScore = score
			bestParent = child
		}
	}
	for _, game := range bestParent {
		game.Results[0].Player.Games = nil
		game.Results[1].Player.Games = nil
	}
	return bestParent
}
