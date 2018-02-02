package tournaments

import (
	"sort"
)

type Pair [2]Player

type PairingParams struct {
	players   []Player
	bracket   int
	origin    bool
	condition func(i int) bool
	cond      string
}

func MakeBye(players []Player, format int) (*Game, []Player) {
	// if even number of player no bye
	if len(players)%2 == 0 {
		return nil, players
	}

	bye := &Game{
		Results: [2]Result{
			Result{
				VictoryPoints:     1,
				ScenarioPoints:    2,
				DestructionPoints: format / 2,
				Bye:               true,
			},
			Result{},
		},
	}
	// stable sort players by victory point incr
	sort.Slice(players, func(i int, j int) bool {
		return players[i].VictoryPoints() < players[j].VictoryPoints()
	})
	for i, p := range players {
		// give a bye to the first player that don't aldready have a bye
		if !p.HadBye() {
			bye.Results[0].PlayerID = p.ID
			players = append(players[:i], players[i+1:]...)
			return bye, players
		}
	}
	// every player already had a bye ? seriously ?
	bye.Results[0].PlayerID = players[0].ID
	players = players[1:]

	//FIXME: maybe we could made this simplier by just giving a bye to the first player

	return bye, players
}

func MakeBrackets(players []Player) (map[int][]Player, []int) {
	brackets := map[int][]Player{}
	for _, p := range players {
		vp := p.VictoryPoints()
		brackets[vp] = append(brackets[vp], p)
	}

	keys := make([]int, 0, len(brackets))
	for k := range brackets {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i int, j int) bool {
		return keys[i] < keys[j]
	})

	for i, vp := range keys {
		if len(brackets[vp])%2 != 0 {
			for j, p := range brackets[vp] {
				if !p.HadSousApp() && p.VictoryPoints() == vp {
					brackets[keys[i+1]] = append(brackets[keys[i+1]], brackets[vp][j])
					brackets[vp] = append(brackets[vp][:j], brackets[vp][j+1:]...)
					break
				}
			}
		}
	}

	return brackets, keys
}

func (t *Tournament) MakePairings(brackets map[int][]Player, keys []int) []Pair {
	// fmt.Println("Make paring")
	pairs := []Pair{}
	for _, vp := range keys {
		players := brackets[vp]
		for len(players) > 1 {
			params := []PairingParams{
				{players, vp, false, func(i int) bool { return i == 1 }, "=1"},
				{players, vp, true, func(i int) bool { return i == 1 }, "=1"},
				{players, vp, true, func(i int) bool { return i > 1 }, ">1"},
				{players, vp, false, func(i int) bool { return i > 1 }, ">1"},
			}
			var pair Pair
			var ok bool
			for _, param := range params {
				if pair, players, ok = t.CreatePair(param); ok {
					pairs = append(pairs, pair)
					break
				}
			}
			if !ok {
				pairs = append(pairs, Pair{players[0], players[1]})
				players = append(players[:0], players[2:]...)
			}
		}
	}
	return pairs
}

func (t *Tournament) CreatePair(params PairingParams) (Pair, []Player, bool) {
	for i, p := range params.players {
		params.players[i].AvailableOpponents = p.GetAvailableOpponents(params.players, params.origin)
	}
	sort.Slice(params.players, func(i int, j int) bool {
		return len(params.players[i].AvailableOpponents) < len(params.players[j].AvailableOpponents)
	})
	var pair Pair
	for i, p := range params.players {
		// log.Debugf("availableOpponents for %s: %+v", p.ID, p.AvailableOpponents)
		if params.condition(len(p.AvailableOpponents)) {
			opponent := t.Players[p.AvailableOpponents[0]]
			pair = Pair{p, opponent}
			params.players = append(params.players[:i], params.players[i+1:]...)
			opponentIndex := getPlayerIndex(opponent, params.players)
			params.players = append(params.players[:opponentIndex], params.players[opponentIndex+1:]...)
			// log.Debugf("Pair %s (%s) vs %s (%s) with params %t, %s", pair[0].ID, pair[0].Origin, pair[1].ID, pair[1].Origin, params.origin, params.cond)
			return pair, params.players, true
		}
	}
	// log.Debugf("No pair with params %t, %s", params.origin, params.cond)
	return pair, params.players, false
}
