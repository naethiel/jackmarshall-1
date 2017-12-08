package tournaments

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	"github.com/chibimi/jackmarshall/tournaments/solver"
)

type Pair [2]*Player

func PairsFromPlayers(players []*Player) []Pair {
	res := []Pair{}
	for i := 0; i < len(players); i++ {
		if i == len(players)-1 {
			break
		}
		res = append(res, Pair{players[i], players[i+1]})
		i++
	}
	return res
}

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
		if p[i].Origin == p[i+1].Origin && p[i].Origin != "" {
			fitness += math.Pow(float64(len(p)), 0)
		}
	}
	return fitness
}

func (p Players) Mutate(randomSwapRate float64) *solver.Individual {
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
