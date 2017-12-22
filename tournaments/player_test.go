package tournaments

import (
	"reflect"
	"testing"
)

func TestAvailableOpponents(t *testing.T) {
	type Case struct {
		lib          string
		playerID     string
		playerOrigin string
		opponents    []string
		players      []Player
		origin       bool
		res          []string
	}

	cases := []Case{
		{
			lib:          "Without origin",
			playerID:     "p1",
			playerOrigin: "Paris",
			opponents:    []string{"p2", "p3"},
			players: []Player{
				{ID: "p1", Origin: "Paris"},
				{ID: "p2", Origin: "Toulon"},
				{ID: "p3", Origin: "Lille"},
				{ID: "p4", Origin: "Paris"},
				{ID: "p5", Origin: "Lyon"},
			},
			origin: false,
			res:    []string{"p4", "p5"},
		},
		{
			lib:          "With origin and all players have origin",
			playerID:     "p1",
			playerOrigin: "Paris",
			opponents:    []string{"p2", "p3"},
			players: []Player{
				{ID: "p1", Origin: "Paris"},
				{ID: "p2", Origin: "Toulon"},
				{ID: "p3", Origin: "Lille"},
				{ID: "p4", Origin: "Paris"},
				{ID: "p5", Origin: "Lyon"},
			},
			origin: true,
			res:    []string{"p5"},
		},
		{
			lib:          "With origin and some players don't have origin",
			playerID:     "p1",
			playerOrigin: "Paris",
			opponents:    []string{"p2", "p3"},
			players: []Player{
				{ID: "p1", Origin: ""},
				{ID: "p2", Origin: "Lyon"},
				{ID: "p3", Origin: "Lille"},
				{ID: "p4", Origin: ""},
				{ID: "p5", Origin: ""},
			},
			origin: true,
			res:    []string{"p4", "p5"},
		},
	}

	for i, c := range cases {
		p := Player{ID: c.playerID, Origin: c.playerOrigin, Oponnent: c.opponents}

		availables := p.GetAvailableOpponents(c.players, c.origin)

		if !reflect.DeepEqual(availables, c.res) {
			t.Log("unexpected result for case", i, "\n\tgot:", availables, "\n\texpected:", c.res)
			t.Fail()
		}
	}
}
