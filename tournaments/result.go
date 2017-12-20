package tournaments

type Result struct {
	PlayerID          string `json:"player"`
	List              string `json:"list"`
	VictoryPoints     int    `json:"victory_points"`
	ScenarioPoints    int    `json:"scenario_points,string"`
	DestructionPoints int    `json:"destruction_points,string"`
	SoS               int    `json:"sos"`
	Bye               bool   `json:"bye"`
	SousApp           bool   `json:"sous_app"`
}

//
// func (t Tournament) getPlayersWithGames() map[string]*Player {
// 	var players = make(map[string]*Player)
// 	for i := range t.Players {
// 		players[t.Players[i].Name] = t.Players[i]
// 	}
//
// 	for _, player := range players {
// 		player.Games = make([]*Game, 0)
// 		for r, round := range t.Rounds {
// 			for g, game := range round.Games {
// 				for _, result := range game.Results {
// 					if result.Player.Name == player.Name {
// 						player.Games = append(player.Games, &t.Rounds[r].Games[g])
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return players
// }
//
// func (p *Player) CumulatedResults() Result {
// 	cumul := Result{}
// 	result := Result{}
// 	for _, game := range p.Games {
// 		if game.Results[0].Player.Name == p.Name {
// 			result = game.Results[0]
// 		} else {
// 			result = game.Results[1]
// 		}
// 		cumul.VictoryPoints += result.VictoryPoints
// 		cumul.ScenarioPoints += result.ScenarioPoints
// 		cumul.DestructionPoints += result.DestructionPoints
// 	}
// 	return cumul
// }
//
// func (t Tournament) GetResults() []*Result {
// 	var results = make([]*Result, 0)
//
// 	players := t.getPlayersWithGames()
//
// 	//calc each player cumulated results
// 	for _, player := range players {
// 		playerResult := player.CumulatedResults()
// 		playerResult.Player = *player
// 		results = append(results, &playerResult)
// 	}
//
// 	//calc SoS
// 	for _, r := range results {
// 		for _, g := range r.Player.Games {
// 			if g.Results[0].Player.Name == "" || g.Results[1].Player.Name == "" || r.Player.Name == "" {
// 				continue
// 			}
// 			if g.Results[0].Player.Name == r.Player.Name {
// 				r.SoS += players[g.Results[1].Player.Name].VictoryPoints()
// 			} else {
// 				r.SoS += players[g.Results[0].Player.Name].VictoryPoints()
// 			}
// 		}
// 	}
//
// 	//clean player's games
// 	for _, r := range results {
// 		r.Player.Games = nil
// 	}
//
// 	return results
// }
