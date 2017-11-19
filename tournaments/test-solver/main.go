package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	. "github.com/chibimi/jackmarshall/tournaments"
	"github.com/chibimi/jackmarshall/tournaments/solver"
)

type Players []*Player

func (p Players) NewIndividual() *solver.Individual {
	players := make([]*Player, len(p))
	copy(players, p)
	for i := 0; i < len(players); i++ {
		j := rand.Intn(len(players))
		players[i], players[j] = players[j], players[i]
	}
	//Sort player list by victory points
	sort.Slice(players, func(i int, j int) bool {
		return players[i].VictoryPoints() > players[j].VictoryPoints()
	})

	return &solver.Individual{
		Fitness: math.MaxInt32,
		Genes:   Players(players),
	}
}

func (p Players) CalcFitness() float64 {
	var fitness float64 = 0
	for i := 0; i < len(p); i += 2 {
		//same match
		if p[i].PlayedAgainst(p[i+1]) {
			fitness += math.Pow(float64(len(p)), float64(len(p)))
		}

		//score
		gap := math.Abs(float64(p[i].VictoryPoints()) - float64(p[i+1].VictoryPoints()))
		if gap != 0.0 {
			fitness += math.Pow(float64(len(p)), gap)
		}

		//origin
		if p[i].Origin == p[i+1].Origin {
			fitness += math.Pow(float64(len(p)), 0)
		}
	}
	return fitness
}

func (p Players) Mutate(randomSwapRate float64) *solver.Individual {
	// i := rand.Intn(len(p) - 1)
	// temp := p[i+1]
	// p[i+1] = p[i]
	// p[i] = temp

	var child = make(Players, len(p))
	copy(child, p)

	r := rand.Float64()
	switch {
	case (r < randomSwapRate):
		i := rand.Intn(len(child) - 1)
		j := rand.Intn(len(child) - 1)
		child[i], child[j] = child[j], child[i]
	default:
		i := rand.Intn(len(child) - 1)
		child[i], child[i+1] = child[i+1], child[i]
	}
	return &solver.Individual{Genes: child}

}

func (p Players) String() string {
	s := ""
	for _, player := range p {
		s = fmt.Sprintf("%s%s, ", s, player.Name)
	}
	return s
}

func main() {
	t := NewTestTournament(60, 32, 32)
	nbRound := 3

	sol := solver.Solver{
		PopulationSize:   10,
		MaxIterations:    10000,
		NumberOfChildren: 10,
		RandomSwapRate:   0.5,
	}

	for i := 0; i < nbRound; i++ {
		t.AddPlayersGames()
		players := t.GetActivePlayers()

		round := Round{
			Number: i,
			Games:  []Game{},
		}

		// fmt.Println("==> Active players")
		// for _, v := range players {
		// 	fmt.Println(v.String())
		// }

		// res := s.Solve(players)
		res2, n := sol.Solve(Players(players))

		fmt.Println(n, res2.Fitness, res2.Genes.CalcFitness())

		pairings := PairsFromPlayers(res2.Genes.(Players))
		fmt.Println(pairings)
		CreateRound(pairings, t.Tables, &round)

		for i, _ := range round.Games {
			round.Games[i].Results[rand.Intn(2)].VictoryPoints = 1
		}

		t.Rounds = append(t.Rounds, round)

	}

	fmt.Println("FINI ! ")
}
