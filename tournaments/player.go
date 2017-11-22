package tournaments

import (
	"fmt"
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
