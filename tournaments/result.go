package tournaments

type Result struct {
	PlayerID          string `json:"player"`
	List              string `json:"list"`
	VictoryPoints     int    `json:"victory_points"`
	ScenarioPoints    int    `json:"scenario_points,string"`
	DestructionPoints int    `json:"destruction_points,string"`
	CasterKill        bool   `json:"caster_kill"`
	SoS               int    `json:"sos"`
	Bye               bool   `json:"bye"`
	SousApp           bool   `json:"sous_app"`
}

type TournamentResult struct {
	PlayerID          string `json:"player"`
	VictoryPoints     int    `json:"victory_points"`
	ScenarioPoints    int    `json:"scenario_points,string"`
	DestructionPoints int    `json:"destruction_points,string"`
	CasterKill        int    `json:"caster_kills"`
	SoS               int    `json:"sos"`
}

func (t Tournament) SetResults() {
	for _, p := range t.Players {
		for _, r := range t.Rounds {
			for _, g := range r.Games {
				var res Result
				if g.Results[0].PlayerID == p.ID {
					res = g.Results[0]
				} else if g.Results[1].PlayerID == p.ID {
					res = g.Results[1]
				}
				player := t.Players[p.ID]
				player.Result.VictoryPoints += res.VictoryPoints
				player.Result.ScenarioPoints += res.ScenarioPoints
				player.Result.DestructionPoints += res.DestructionPoints
				if res.CasterKill {
					player.Result.CasterKill += 1
				}
				t.Players[p.ID] = player
			}
		}
	}

	//calc SoS
	for _, r := range t.Rounds {
		for _, g := range r.Games {
			p0 := t.Players[g.Results[0].PlayerID]
			p1 := t.Players[g.Results[1].PlayerID]
			p0.Result.SoS += p1.Result.VictoryPoints
			p1.Result.SoS += p0.Result.VictoryPoints
			t.Players[p0.ID] = p0
			t.Players[p1.ID] = p1
		}
	}
}
