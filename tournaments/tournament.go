package tournaments

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/chibimi/jackmarshall/tournaments/solver"
	"gopkg.in/mgo.v2/bson"
)

type Tournament struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty" update:"nonzero"`
	Owner     int64         `json:"owner" bson:"owner" create:"nonzero" update:"nonzero"`
	Name      string        `json:"name" create:"nonzero" update:"nonzero"`
	Format    int           `json:"format" create:"nonzero" update:"nonzero"`
	Slots     int           `json:"slots" create:"nonzero" update:"nonzero"`
	FeeAmount float64       `json:"fee_amount"`
	Date      time.Time     `json:"date" create:"nonzero" update:"nonzero"`
	Players   []*Player     `json:"players"`
	Tables    []Table       `json:"tables"`
	Rounds    []Round       `json:"rounds" create:"max=0"`
}

func NewTournament() *Tournament {
	return &Tournament{
		Players: []*Player{},
		Tables:  []Table{},
		Rounds:  []Round{},
	}
}

func (t *Tournament) GetActivePlayers() []*Player {
	res := []*Player{}
	for i := range t.Players {
		if t.Players[i].Leave != true {
			res = append(res, t.Players[i])
		}
	}
	return res
}

func (t *Tournament) AddPlayersGames() {
	games := map[string][]*Game{}

	for _, r := range t.Rounds {
		for i, g := range r.Games {
			for _, res := range g.Results {
				if v, ok := games[res.Player.ID]; ok {
					games[res.Player.ID] = append(v, &r.Games[i])
				} else {
					games[res.Player.ID] = []*Game{&r.Games[i]}
				}
			}
		}
	}

	for _, p := range t.Players {
		p.Games = games[p.ID]
	}
}

func (t *Tournament) CreateNextRound() {
	//Init solver
	s := solver.Solver{
		PopulationSize:   10,
		MaxIterations:    10000,
		NumberOfChildren: 10,
		RandomSwapRate:   0.5,
	}
	//Create pairing
	t.AddPlayersGames()
	players := Players(t.GetActivePlayers())
	pairing, i := s.Solve(players)
	fmt.Printf("Pairing done in %d itarations with a fitness score of %.0f\n", i, pairing.Fitness)

	pairs := PairsFromPlayers(pairing.Genes.(Players))
	assignements, j := s.Solve(Assignements{Pairs: pairs, Tables: t.Tables})
	fmt.Printf("Assignements done in %d itarations with a fitness score of %.0f\n", j, assignements.Fitness)
}

func NewTestTournament(nbPlayer, nbTable, nbScenario int) *Tournament {
	t := NewTournament()

	var origins = []string{"whag", "loin", "uchro", "ludo", "uchro", "usa"}

	for i := 0; i < nbPlayer; i++ {
		t.Players = append(t.Players, &Player{
			ID:     "player" + fmt.Sprintf("%d", i),
			Name:   "player" + fmt.Sprintf("%d", i),
			Origin: origins[rand.Intn(len(origins))],
		})
	}

	for i := 0; i < nbTable; i++ {
		t.Tables = append(t.Tables, Table{
			ID:       "table" + fmt.Sprintf("%d", i),
			Name:     "table" + fmt.Sprintf("%d", i),
			Scenario: "scenario" + fmt.Sprintf("%d", i%nbScenario),
		})
	}

	return t
}
