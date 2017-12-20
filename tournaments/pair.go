package tournaments

import (
	"fmt"
	"objenious/lib/log"
	"sort"
)

type Pair [2]Player

type PairingParams struct {
	// p         Player
	players   []Player
	bracket   int
	origin    bool
	condition func(i int) bool
	cond      string
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
		log.Debugf("availableOpponents for %s: %+v", p.ID, p.AvailableOpponents)
		if params.condition(len(p.AvailableOpponents)) {
			opponent := t.Players[p.AvailableOpponents[0]]
			pair = Pair{p, opponent}
			params.players = append(params.players[:i], params.players[i+1:]...)
			opponentIndex := getPlayerIndex(opponent, params.players)
			params.players = append(params.players[:opponentIndex], params.players[opponentIndex+1:]...)
			log.Debugf("Pair %s (%s) vs %s (%s) with params %t, %s", pair[0].ID, pair[0].Origin, pair[1].ID, pair[1].Origin, params.origin, params.cond)
			return pair, params.players, true
		}
	}
	log.Debugf("No pair with params %t, %s", params.origin, params.cond)
	return pair, params.players, false
}

func printPlayers(players []Player) {
	s := "Players: "
	for _, p := range players {
		s = fmt.Sprintf("%s, %s", s, p.ID)
	}
	fmt.Println(s)
}

//

// //
// func PairsFromPlayers(players []*Player) []Pair {
// 	res := []Pair{}
// 	for i := 0; i < len(players); i++ {
// 		if i == len(players)-1 {
// 			break
// 		}
// 		res = append(res, Pair{players[i], players[i+1]})
// 		i++
// 	}
// 	return res
// }
//
// type Players []*Player
//
// func (p Players) NewIndividual() *solver.Individual {
// 	players := make([]*Player, len(p))
// 	copy(players, p)
// 	for i := 0; i < len(players); i++ {
// 		j := rand.Intn(len(players))
// 		players[i], players[j] = players[j], players[i]
// 	}
// 	//Sort player list by victory points
// 	sort.Slice(players, func(i int, j int) bool {
// 		return players[i].VictoryPoints() > players[j].VictoryPoints()
// 	})
//
// 	return &solver.Individual{
// 		Fitness: math.MaxInt32,
// 		Genes:   Players(players),
// 	}
// }
//
// func (p Players) CalcFitness() float64 {
// 	var fitness float64 = 0
// 	for i := 0; i < len(p); i += 2 {
// 		//same match
// 		if p[i].PlayedAgainst(p[i+1]) {
// 			fitness += math.Pow(float64(len(p)), float64(len(p)))
// 		}
//
// 		//score
// 		gap := math.Abs(float64(p[i].VictoryPoints()) - float64(p[i+1].VictoryPoints()))
// 		if gap != 0.0 {
// 			fitness += math.Pow(float64(len(p)), gap)
// 		}
//
// 		//origin
// 		if p[i].Origin == p[i+1].Origin && p[i].Origin != "" {
// 			fitness += 1
// 		}
// 	}
// 	return fitness
// }
//
// func (p Players) Mutate(randomSwapRate float64) *solver.Individual {
// 	var child = make(Players, len(p))
// 	copy(child, p)
//
// 	r := rand.Float64()
// 	switch {
// 	case (r < randomSwapRate):
// 		i := rand.Intn(len(child) - 1)
// 		j := rand.Intn(len(child) - 1)
// 		child[i], child[j] = child[j], child[i]
// 	default:
// 		i := rand.Intn(len(child) - 1)
// 		child[i], child[i+1] = child[i+1], child[i]
// 	}
// 	return &solver.Individual{Genes: child}
// }
//
// func (p Players) String() string {
// 	s := ""
// 	for _, player := range p {
// 		s = fmt.Sprintf("%s%s, ", s, player.Name)
// 	}
// 	return s
// }
