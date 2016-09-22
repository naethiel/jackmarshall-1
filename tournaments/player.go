package main

type Player struct {
	Name     string    `json:"name"`
	Faction  string    `json:"faction"`
	PayedFee bool      `json:"payed_fee"`
	Lists    [2]string `json:"lists"`
	// Leave    bool     `json:"leave"`
	// Games    []*Game
}

//
// func (p *Player) String() string {
// 	return p.Name
// }
//
// func (p Player) Score() int {
// 	var score = 0
// 	for _, game := range p.Games {
// 		var result = game.Results[0]
// 		if result.Player != p.Name {
// 			result = game.Results[1]
// 		}
//
// 		score += result.VictoryPoints
// 	}
// 	return score
// }
//
// func (p Player) hadBye() bool {
// 	for _, game := range p.Games {
// 		fmt.Println(game.Results)
// 		for _, result := range game.Results {
// 			if result.Player == p.Name && result.Buy == true {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
//
// func (p Player) PlayedAgainst(o string) bool {
// 	for _, game := range p.Games {
// 		for _, result := range game.Results {
// 			if result.Player == o {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
//
// func (p Player) PlayedOn(t Table) bool {
// 	for _, game := range p.Games {
// 		if game.Table.Name == t.Name {
// 			return true
// 		}
// 	}
// 	return false
// }
//
// func (p *Player) AddGame(g *Game) {
// 	p.Games = append(p.Games, g)
// }
