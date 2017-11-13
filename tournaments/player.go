package tournaments

import "fmt"

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

func (p *Player) String() string {
	return fmt.Sprintf("%s (%d/%d)", p.Name, p.VictoryPoints(), len(p.Games))
}
