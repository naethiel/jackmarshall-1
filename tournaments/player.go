package tournaments

type Player struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	Origin             string   `json:"origin"`
	Faction            string   `json:"faction"`
	PayedFee           bool     `json:"payed_fee"`
	Leave              bool     `json:"leave"`
	Games              []Game   `json:"-"`
	Oponnent           []string `json:"-"`
	Tables             []Table  `json:"-"`
	AvailableOpponents []string `json:"-"`
	Result             Result   `json:"result"`
}

func (p *Player) VictoryPoints() int {
	score := 0
	for _, g := range p.Games {
		res := g.Results[0]
		if res.PlayerID != p.ID {
			res = g.Results[1]
		}
		score += res.VictoryPoints
	}
	return score
}

func (p *Player) HadBye() bool {
	for _, g := range p.Games {
		for _, res := range g.Results {
			if res.PlayerID == p.ID && res.Bye == true {
				return true
			}
		}
	}
	return false
}

func (p *Player) HadSousApp() bool {
	for _, g := range p.Games {
		for _, res := range g.Results {
			if res.PlayerID == p.ID && res.SousApp == true {
				return true
			}
		}
	}
	return false
}

func (p *Player) PlayedAgainst(opponent string) bool {
	for _, o := range p.Oponnent {
		if o == opponent {
			return true
		}
	}
	return false
}

func (p *Player) NbPlayedTable(tableID string) int {
	res := 0
	for _, t := range p.Tables {
		if t.ID == tableID {
			res++
		}
	}
	return res
}

func (p *Player) NbPlayedScenario(scenario string) int {
	res := 0
	for _, t := range p.Tables {
		if t.Scenario == scenario {
			res++
		}
	}
	return res
}

func (p *Player) GetAvailableOpponents(players []Player, origin bool) []string {
	res := []string{}
	for _, opponent := range players {
		if opponent.ID == p.ID || p.PlayedAgainst(opponent.ID) {
			continue
		}
		if origin && p.Origin != "" && p.Origin == opponent.Origin {
			continue
		}
		res = append(res, opponent.ID)
	}
	return res
}
