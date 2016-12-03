package main

type Player struct {
	Name     string `json:"name"`
	Faction  string `json:"faction"`
	PayedFee bool   `json:"payed_fee"`
	Lists    []List `json:"lists"`
	Leave    bool   `json:"leave"`
	Games    []*Game
}

func (p Player) VictoryPoints() int {
	var score = 0
	for _, game := range p.Games {
		var result = game.Results[0]
		if result.Player.Name != p.Name {
			result = game.Results[1]
		}
		score += result.VictoryPoints
	}
	return score
}

func (p Player) CumulatedResults() Result {
	cumul := Result{}
	result := Result{}
	for _, game := range p.Games {
		if game.Results[0].Player.Name == p.Name {
			result = game.Results[0]
		} else {
			result = game.Results[1]
		}
		cumul.VictoryPoints += result.VictoryPoints
		cumul.ScenarioPoints += result.ScenarioPoints
		cumul.DestructionPoints += result.DestructionPoints
	}
	return cumul
}

func (p Player) hadBye() bool {
	for _, game := range p.Games {
		//fmt.Println(game.Results)
		for _, result := range game.Results {
			if result.Player.Name == p.Name && result.Bye == true {
				return true
			}
		}
	}
	return false
}

func (p Player) PlayedAgainst(o string) bool {
	for _, game := range p.Games {
		for _, result := range game.Results {
			if result.Player.Name == o {
				return true
			}
		}
	}
	return false
}

func (p Player) PlayedOn(scenario string) bool {
	for _, game := range p.Games {
		if game.Table.Scenario == scenario {
			return true
		}
	}
	return false
}
