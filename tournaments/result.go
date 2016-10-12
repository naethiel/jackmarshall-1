package main

import "encoding/json"

//Result represents the result of a player for a game
type Result struct {
	Player            Player `json:"player"`
	List              int    `json:"list"`
	VictoryPoints     int    `json:"victory_points"`
	ScenarioPoints    int    `json:"scenario_points"`
	DestructionPoints int    `json:"destruction_points"`
	Bye               bool   `json:"bye"`
}

func getCumulatedResult(tournament Tournament) []Result {
	var res = make([]Result, 0)
	var players = make([]*Player, len(tournament.Players))
	json.MarshalIndent(tournament, "", "\t")
	for i := range tournament.Players {
		players[i] = &tournament.Players[i]
	}

	for _, player := range players {
		player.Games = make([]*Game, 0)
		for r, round := range tournament.Rounds {
			for g, game := range round.Games {
				for _, result := range game.Results {
					if result.Player.Name == player.Name {
						player.Games = append(player.Games, &tournament.Rounds[r].Games[g])
					}
				}
			}
		}

	}

	for _, player := range players {
		cumul := player.CumulatedResults()
		player.Games = nil
		cumul.Player = *player
		res = append(res, cumul)
	}
	return res
}
