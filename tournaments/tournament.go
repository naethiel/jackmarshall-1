package tournaments

import (
	"fmt"
	"time"

	"github.com/chibimi/jackmarshall/tournaments/solver"
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

func (t *Tournament) SetPreviousGamesData() {
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
	t.SetPreviousGamesData()
	players := t.GetActivePlayers()
	suffle(players)

	bye, players := MakeBye(players, t.Format)

	brackets, keys := MakeBrackets(players)

	pairs := t.MakePairings(brackets, keys)

	// round := Round{}
	// for _, p := range pairs {
	// 	g := Game{
	// 		Results: [2]Result{
	// 			{PlayerID: p[0].ID},
	// 			{PlayerID: p[1].ID},
	// 		},
	// 	}
	// 	round.Games = append(round.Games, g)
	// }
	s := solver.Solver{
		PopulationSize:   10,
		MaxIterations:    10000,
		NumberOfChildren: 10,
		RandomSwapRate:   0.5,
	}
	//Assign tables
	tables := []Table{}
	for _, v := range t.Tables {
		tables = append(tables, v)
	}
	assignements, j := s.Solve(Assignements{Pairs: pairs, Tables: tables})
	fmt.Printf("Assignements done in %d itarations with a fitness score of %.0f\n", j, assignements.Fitness)
	round := RoundFromAssignaments(assignements.Genes.(Assignements))
	if bye != nil {
		round.Games = append(round.Games, *bye)
	}
	round.Number = len(t.Rounds)
	return round
}
