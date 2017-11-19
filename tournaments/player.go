package tournaments

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	"github.com/chibimi/jackmarshall/tournaments/solver"
)

type Player struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Origin   string `json:"origin"`
	Faction  string `json:"faction"`
	PayedFee bool   `json:"payed_fee"`
	Lists    []List `json:"lists"`
	Leave    bool   `json:"leave"`
	Games    []*Game
}

func (p *Player) VictoryPoints() int {
	score := 0
	for _, g := range p.Games {
		res := g.Results[0]
		if res.Player.ID != p.ID {
			res = g.Results[1]
		}
		score += res.VictoryPoints
	}
	return score
}

func (p *Player) HadBye() bool {
	for _, g := range p.Games {
		for _, res := range g.Results {
			if res.Player.ID == p.ID && res.Bye == true {
				return true
			}
		}
	}
	return false
}

func (p *Player) PlayedAgainst(o *Player) bool {
	for _, game := range p.Games {
		for _, result := range game.Results {
			if result.Player.ID == o.ID {
				return true
			}
		}
	}
	return false
}

func (p *Player) PlayedScenario(scenario string) bool {
	for _, game := range p.Games {
		if game.Table.Scenario == scenario {
			return true
		}
	}
	return false
}

func (p *Player) PlayedTable(table Table) bool {
	for _, game := range p.Games {
		if game.Table.ID == table.ID {
			return true
		}
	}
	return false
}

func (p *Player) NbPlayedScenario(scenario string) int {
	res := 0
	for _, game := range p.Games {
		if game.Table.Scenario == scenario {
			res++
		}
	}
	return res
}

func (p *Player) NbPlayedTable(table Table) int {
	res := 0
	for _, game := range p.Games {
		if game.Table.ID == table.ID {
			res++
		}
	}
	return res
}

func (p *Player) String() string {
	return fmt.Sprintf("%s %s (%d)", p.Name, p.Origin, p.VictoryPoints())
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
