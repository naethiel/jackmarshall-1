package main

import "time"

type Tournament struct {
	Id        string    `json:"id" update:"nonzero"`
	Name      string    `json:"name" create:"nonzero" update:"nonzero"`
	Format    int       `json:"format" create:"nonzero" update:"nonzero"`
	Slots     int       `json:"slots" create:"nonzero" update:"nonzero"`
	FeeAmount float64   `json:"fee_amount"`
	Date      time.Time `json:"date" create:"nonzero" update:"nonzero"`
	Players   []Player  `json:"players"`
	Tables    []Table   `json:"tables"`
	Rounds    []Round   `json:"rounds" create:"max=0"`
}

func NewTournament() *Tournament {
	t := new(Tournament)
	t.Players = []Player{}
	t.Tables = []Table{}
	t.Rounds = []Round{}
	return t
}

func (t Tournament) getPlayersWithGames() map[string]*Player {
	var players = make(map[string]*Player)
	for i := range t.Players {
		players[t.Players[i].Name] = &t.Players[i]
	}

	for _, player := range players {
		player.Games = make([]*Game, 0)
		for r, round := range t.Rounds {
			for g, game := range round.Games {
				for _, result := range game.Results {
					if result.Player.Name == player.Name {
						player.Games = append(player.Games, &t.Rounds[r].Games[g])
					}
				}
			}
		}
	}
	return players
}

func (t Tournament) getResults() []*Result {
	var results = make([]*Result, 0)

	players := t.getPlayersWithGames()

	//calc each player cumulated results
	for _, player := range players {
		playerResult := player.CumulatedResults()
		playerResult.Player = *player
		results = append(results, &playerResult)
	}

	//calc SoS
	for _, r := range results {
		for _, g := range r.Player.Games {
			if g.Results[0].Player.Name == r.Player.Name {
				r.SoS += players[g.Results[1].Player.Name].VictoryPoints()
			} else {
				r.SoS += players[g.Results[0].Player.Name].VictoryPoints()
			}
		}
	}

	//clean player's games
	for _, r := range results {
		r.Player.Games = nil
	}

	return results
}
