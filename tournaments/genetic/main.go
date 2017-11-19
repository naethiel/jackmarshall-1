package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"

	"github.com/chibimi/jackmarshall/tournaments"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// type Player struct {
// 	Name   string
// 	Score  int
// 	Origin string
// }

// func createPlayers(nbPlayer int) []Player {
// 	players := make([]Player, nbPlayer)
// 	for i := 0; i < nbPlayer; i++ {
// 		players[i] = Player{
// 			fmt.Sprintf("player %s", i),
// 			2,
// 			origins[rand.Intn(len(origins))],
// 		}
// 	}
// 	return players
// }

func main() {
	// var players = []Player{
	// 	{"pim", 2, "whag"},
	// 	{"juju", 2, "ludo"},
	// 	{"pa", 2, "ludo"},
	// 	{"tza", 2, "lwin"},
	// 	{"julien", 1, "usa"},
	// 	{"plazma", 1, "uchro"},
	// 	{"tec", 1, "ludo"},
	// 	{"largoo", 1, "whag"},
	// 	{"lelith", 1, "lwin"},
	// 	{"ygemethor", 1, "lwin"},
	// 	{"tarasputin", 1, "lwin"},
	// 	{"flo", 1, "lwin"},
	// 	{"yvanass", 0, "uchro"},
	// 	{"haze", 0, "uchro"},
	// 	{"lugmi", 0, "uchro"},
	// 	{"felix", 0, "ludo"},
	// }
	//
	// players = pair(players, Options{
	// 	PopulationSize:   10,
	// 	MaxIterations:    1000,
	// 	NumberOfChildren: 10,
	// 	RandomSwapRate:   0.5,
	// })
	// fmt.Println(fitness(players), players)

	t := tournaments.NewTestTournament(60, 32, 32)
	nbRound := 3

	for i := 0; i < nbRound; i++ {
		t.AddPlayersGames()

		players := t.GetActivePlayers()

		round := tournaments.Round{
			Number: i,
			Games:  []tournaments.Game{},
		}

		// res := s.Solve(players)
		pairs, ite := pair(players, Options{
			PopulationSize:   10,
			MaxIterations:    10000,
			NumberOfChildren: 10,
			RandomSwapRate:   0.5,
		})
		fmt.Println(ite, fitness(pairs))

		pairings := tournaments.PairsFromPlayers(pairs)
		tournaments.CreateRound(pairings, t.Tables, &round)

		for i, _ := range round.Games {
			round.Games[i].Results[rand.Intn(2)].VictoryPoints = 1
		}

		t.Rounds = append(t.Rounds, round)
	}

}

type Options struct {
	PopulationSize   int
	MaxIterations    int
	NumberOfChildren int
	RandomSwapRate   float64
}

func pair(players []*tournaments.Player, options Options) ([]*tournaments.Player, int) {
	var population [][]*tournaments.Player

	for i := 0; i < options.PopulationSize; i++ {
		population = append(population, generate(players))
	}
	s := 0
	for ; s < options.MaxIterations && fitness(population[0]) != 0; s++ {
		for i := 0; i < options.PopulationSize; i++ {
			population = append(population, mutate(population[i], options))
		}

		sort.Slice(population, func(i, j int) bool {
			return fitness(population[i]) < fitness(population[j])
		})

		population = population[:options.PopulationSize]
		// fmt.Println(s, fitness(population[0]), population[0])
	}

	return population[0], s
}

func generate(players []*tournaments.Player) []*tournaments.Player {
	var out = make([]*tournaments.Player, len(players))
	copy(out, players)
	for i := 0; i < len(out); i++ {
		j := rand.Intn(len(out))
		out[i], out[j] = out[j], out[i]
	}
	return out
}

func mutate(parent []*tournaments.Player, options Options) []*tournaments.Player {
	var child = make([]*tournaments.Player, len(parent))
	copy(child, parent)

	p := rand.Float64()
	switch {
	case (p < options.RandomSwapRate):
		i := rand.Intn(len(child) - 1)
		j := rand.Intn(len(child) - 1)
		child[i], child[j] = child[j], child[i]
	default:
		i := rand.Intn(len(child) - 1)
		child[i], child[i+1] = child[i+1], child[i]
	}
	return child
}

func fitness(individual []*tournaments.Player) float64 {
	var fitness float64 = 0
	for i := 0; i < len(individual); i += 2 {
		//same match
		if individual[i].PlayedAgainst(individual[i+1]) {
			fitness += math.Pow(float64(len(individual)), float64(len(individual)))
		}

		//score
		gap := math.Abs(float64(individual[i].VictoryPoints()) - float64(individual[i+1].VictoryPoints()))
		if gap != 0.0 {
			fitness += math.Pow(float64(len(individual)), gap)
		}

		//origin
		if individual[i].Origin == individual[i+1].Origin {
			fitness += math.Pow(float64(len(individual)), 0)
		}
	}
	return fitness
}
