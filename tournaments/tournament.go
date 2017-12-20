package tournaments

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Tournament struct {
	ID      bson.ObjectId     `json:"id" bson:"_id,omitempty" update:"nonzero"`
	Owner   int64             `json:"owner" bson:"owner" create:"nonzero" update:"nonzero"`
	Name    string            `json:"name" create:"nonzero" update:"nonzero"`
	Format  int               `json:"format" create:"nonzero" update:"nonzero"`
	Slots   int               `json:"slots" create:"nonzero" update:"nonzero"`
	Date    time.Time         `json:"date" create:"nonzero" update:"nonzero"`
	Players map[string]Player `json:"players"`
	Tables  map[string]Table  `json:"tables"`
	Rounds  []Round           `json:"rounds" create:"max=0"`
}

func NewTournament() *Tournament {
	return &Tournament{
		Players: map[string]Player{},
		Tables:  map[string]Table{},
		Rounds:  []Round{},
	}
}

func (t *Tournament) GetActivePlayers() []Player {
	res := []Player{}
	for _, v := range t.Players {
		if v.Leave != true {
			res = append(res, v)
		}
	}
	return res
}

func (t *Tournament) SetPlayersGamesData() {
	games := map[string][]Game{}
	tables := map[string][]Table{}
	opponents := map[string][]string{}
	for _, r := range t.Rounds {
		for _, g := range r.Games {
			for i, res := range g.Results {
				games[res.PlayerID] = append(games[res.PlayerID], g)
				opponents[res.PlayerID] = append(opponents[res.PlayerID], g.Results[(i+1)%2].PlayerID)
				tables[res.PlayerID] = append(tables[res.PlayerID], t.Tables[g.TableID])
			}
		}
	}
	for i, p := range t.Players {
		player := t.Players[i]
		player.Games = games[p.ID]
		player.Tables = tables[p.ID]
		player.Oponnent = opponents[p.ID]
		t.Players[i] = player
	}
}

func (t *Tournament) GetNextRound() Round {
	t.SetPlayersGamesData()
	players := t.GetActivePlayers()
	suffle(players)

	var bye Game
	if len(players)%2 != 0 {
		bye, players = setBye(players, t.Format)
	}

	pairs := []Pair{}
	brackets, keys := makeBrackets(players)
	for _, vp := range keys {
		players := brackets[vp]
		fmt.Println("Braket ", vp, len(players))
		for len(players) >= 2 {
			fmt.Println("Len players", len(players))
			params := []PairingParams{
				{players, vp, false, func(i int) bool { return i == 1 }, "=1"},
				{players, vp, true, func(i int) bool { return i == 1 }, "=1"},
				{players, vp, true, func(i int) bool { return i > 1 }, ">1"},
				{players, vp, false, func(i int) bool { return i > 1 }, ">1"},
			}
			var pair Pair
			var ok bool
			for _, param := range params {
				pair, players, ok = t.CreatePair(param)
				if ok {
					pairs = append(pairs, pair)
					break
				}
			}
			if !ok {
				fmt.Println("Fuckers", players[0].ID, players[1].ID)
				pairs = append(pairs, Pair{players[0], players[1]})
				players = append(players[:0], players[2:]...)
			}

		}
	}
	round := Round{}
	for _, p := range pairs {
		g := Game{
			Results: [2]Result{
				{PlayerID: p[0].ID},
				{PlayerID: p[1].ID},
			},
		}
		round.Games = append(round.Games, g)
	}
	// s := solver.Solver{
	// 	PopulationSize:   10,
	// 	MaxIterations:    10000,
	// 	NumberOfChildren: 10,
	// 	RandomSwapRate:   0.5,
	// }
	// //Assign tables
	// tables := []Table{}
	// for _, v := range t.Tables {
	// 	tables = append(tables, v)
	// }
	// assignements, j := s.Solve(Assignements{Pairs: pairs, Tables: tables})
	// fmt.Printf("Assignements done in %d itarations with a fitness score of %.0f\n", j, assignements.Fitness)
	// round := RoundFromAssignaments(assignements.Genes.(Assignements))
	if bye.Results[0].PlayerID != "" {
		round.Games = append(round.Games, bye)
	}
	round.Number = len(t.Rounds)
	return round
}

func suffle(t []Player) {
	for i := 0; i < len(t); i++ {
		j := rand.Intn(len(t))
		t[i], t[j] = t[j], t[i]
	}
}

func setBye(players []Player, format int) (Game, []Player) {
	bye := Game{
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
	sort.Slice(players, func(i int, j int) bool {
		return players[i].VictoryPoints() < players[j].VictoryPoints()
	})
	for i, p := range players {
		if !p.HadBye() {
			bye.Results[0].PlayerID = p.ID
			players = append(players[:i], players[i+1:]...)
			break
		}
	}
	return bye, players
}

func makeBrackets(players []Player) (map[int][]Player, []int) {
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
		return keys[i] > keys[j]
	})

	for i, vp := range keys {
		if len(brackets[vp])%2 != 0 {
			for _, p := range players {
				if !p.HadSousApp() && p.VictoryPoints() == vp {
					brackets[keys[i+1]] = append(brackets[keys[i+1]], p)
					brackets[vp] = append(brackets[vp][:i], brackets[vp][i+1:]...)
					printBracket(brackets)
					break
				}
			}
		}

	}
	return brackets, keys
}

func printBracket(b map[int][]Player) {

	for k, v := range b {
		s := fmt.Sprintf("Bracket %d: ", k)
		for _, p := range v {
			s = fmt.Sprintf("%s, %s", s, p.ID)
		}
		fmt.Println(s)

	}
}

func getPlayerIndex(p Player, tab []Player) int {
	for i, v := range tab {
		if v.ID == p.ID {
			return i
		}
	}
	return -1
}

func NewTestTournament(nbPlayer, nbTable, nbScenario int) *Tournament {
	t := NewTournament()

	var origins = []string{"whag", "loin", "uchro", "ludo", "uchro", ""}
	var factions = []string{"cygnar", "cryx", "legion", "skorne", "trollbloods", ""}

	for i := 0; i < nbPlayer; i++ {
		p := Player{
			ID:      "player" + fmt.Sprintf("%d", i),
			Name:    "player" + fmt.Sprintf("%d", i),
			Origin:  origins[rand.Intn(len(origins))],
			Faction: factions[rand.Intn(len(factions))],
		}
		t.Players[p.ID] = p
	}

	for i := 0; i < nbTable; i++ {
		table := Table{
			ID:       "table" + fmt.Sprintf("%d", i),
			Name:     "table" + fmt.Sprintf("%d", i),
			Scenario: "scenario" + fmt.Sprintf("%d", i%nbScenario),
		}
		t.Tables[table.ID] = table
	}
	t.Date = time.Now().Add(time.Hour)

	return t
}
func (t *Tournament) CheckRound(index int) (int, int, int, int) {
	round := t.Rounds[index]

	pairing := []string{}
	origin := []string{}
	sousApp := []string{}
	bye := []string{}

	for _, g := range round.Games {
		p0 := t.Players[g.Results[0].PlayerID]
		p1 := t.Players[g.Results[1].PlayerID]
		if p0.PlayedAgainst(p1.ID) {
			pairing = append(pairing, fmt.Sprintf("%s vs %s", p0.ID, p1.ID))
		}
		if p0.Origin != "" && p0.Origin == p1.Origin {
			origin = append(origin, fmt.Sprintf("%s vs %s (%s)", p0.ID, p1.ID, p0.Origin))
		}
		if g.Results[0].Bye && p0.HadBye() {
			bye = append(bye, p0.ID)
		}
		if g.Results[0].SousApp && p0.HadSousApp() {
			sousApp = append(sousApp, p0.ID)
		}
		if g.Results[1].SousApp && p1.HadSousApp() {
			sousApp = append(sousApp, p1.ID)
		}

	}

	fmt.Println("Pairing (", len(pairing), "):", pairing)
	fmt.Println("Origin (", len(origin), "):", origin)
	fmt.Println("Bye (", len(bye), "):", bye)
	fmt.Println("SousApp (", len(sousApp), "):", sousApp)

	return len(pairing), len(origin), len(bye), len(sousApp)

}
