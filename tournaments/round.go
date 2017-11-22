package tournaments

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/chibimi/jackmarshall/tournaments/solver"
)

type Round struct {
	Number int    `json:"number"`
	Games  []Game `json:"games"`
}

func (r *Round) String() (s string) {
	for i, game := range r.Games {
		s += fmt.Sprintf("GAME %d :\t%s  %s\t", i, game.Table.Name, game.Table.Scenario)
		s += fmt.Sprintf("%s (%d) vs %s (%d)\n", game.Results[0].Player.Name, game.Results[0].VictoryPoints, game.Results[1].Player.Name, game.Results[1].VictoryPoints)
	}
	return
}

type Assignements struct {
	Pairs  []Pair
	Tables []Table
}

func (a Assignements) NewIndividual() *solver.Individual {
	pairs := make([]Pair, len(a.Pairs))
	copy(pairs, a.Pairs)
	tables := make([]Table, len(a.Tables))
	copy(tables, a.Tables)
	for i := 0; i < len(tables); i++ {
		j := rand.Intn(len(tables))
		tables[i], tables[j] = tables[j], tables[i]
	}
	return &solver.Individual{
		Fitness: math.MaxInt32,
		Genes: Assignements{
			Pairs:  pairs,
			Tables: tables,
		},
	}
}

func (a Assignements) CalcFitness() float64 {
	var fitness float64 = 0

	for i := 0; i < len(a.Pairs); i++ {
		for _, p := range a.Pairs[i] {
			nbTable := p.NbPlayedTable(a.Tables[i])
			nbScenario := p.NbPlayedScenario(a.Tables[i].Scenario)
			if nbScenario != 0 {
				fitness += math.Pow(float64(len(a.Pairs)), float64(nbScenario))
			}
			if nbTable != 0 {
				fitness += float64(nbTable) * math.Pow(float64(len(a.Pairs)), float64(len(a.Pairs)))
			}
		}
	}

	return fitness
}

func (a Assignements) Mutate(randomSwapRate float64) *solver.Individual {
	pairs := make([]Pair, len(a.Pairs))
	copy(pairs, a.Pairs)
	tables := make([]Table, len(a.Tables))
	copy(tables, a.Tables)

	i := rand.Intn(len(tables) - 1)
	j := rand.Intn(len(tables) - 1)
	tables[i], tables[j] = tables[j], tables[i]

	return &solver.Individual{
		Genes: Assignements{
			Pairs:  pairs,
			Tables: tables,
		},
	}
}

func (a Assignements) String() string {
	s := ""
	for i := 0; i < len(a.Pairs); i++ {
		s = fmt.Sprintf("%s%s (%s) : %s (%d, %s) vs %s (%d, %s)\n", s, a.Tables[i].Name, a.Tables[i].Scenario, a.Pairs[i][0].Name, a.Pairs[i][0].VictoryPoints(), a.Pairs[i][0].Origin, a.Pairs[i][1].Name, a.Pairs[i][1].VictoryPoints(), a.Pairs[i][1].Origin)
	}
	return s
}

func RoundFromAssignaments(a Assignements) Round {
	games := []Game{}
	for i := 0; i < len(a.Pairs); i++ {
		games = append(games, Game{
			Table:   a.Tables[i],
			Results: [2]Result{{Player: *a.Pairs[i][0]}, {Player: *a.Pairs[i][1]}},
		})
	}
	return Round{
		Games: games,
	}
}
