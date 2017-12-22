package tournaments

import (
	"reflect"
	"testing"
)

func TestMakeBye(t *testing.T) {

	type Case struct {
		lib      string
		players  []Player
		format   int
		bye      *Game
		lenAfter int
	}

	cases := []Case{
		{
			lib: "p1 has less victory point than p2 ans p3",
			players: []Player{
				{
					ID:    "p1",
					Games: []Game{makeGameResult("t1", "p1", "p2", 0, 1)},
				},
				{
					ID:    "p2",
					Games: []Game{makeGameResult("t1", "p1", "p2", 0, 1)},
				},
				{
					ID:    "p3",
					Games: []Game{makeGameResult("", "p3", "", 1, 0)},
				},
			},
			format: 100,
			bye: &Game{
				TableID: "",
				Results: [2]Result{
					{PlayerID: "p1", VictoryPoints: 1, ScenarioPoints: 2, DestructionPoints: 50, Bye: true},
					{},
				},
			},
			lenAfter: 2,
		},
		{
			lib: "p1 has less victory point than p2 ans p3 but already had a bye",
			players: []Player{
				{
					ID: "p1",
					Games: []Game{
						{
							TableID: "t1",
							Results: [2]Result{
								{PlayerID: "p1", VictoryPoints: 0, Bye: true},
								{PlayerID: "p2", VictoryPoints: 1},
							},
						},
					},
				},
				{
					ID:    "p2",
					Games: []Game{makeGameResult("t1", "p1", "p2", 0, 1)},
				},
				{
					ID:    "p3",
					Games: []Game{makeGameResult("", "p3", "", 2, 0)},
				},
			},
			format: 75,
			bye: &Game{
				TableID: "",
				Results: [2]Result{
					{PlayerID: "p2", VictoryPoints: 1, ScenarioPoints: 2, DestructionPoints: 37, Bye: true},
					{},
				},
			},
			lenAfter: 2,
		},
	}

	for i, c := range cases {
		bye, players := MakeBye(c.players, c.format)
		if !reflect.DeepEqual(bye, c.bye) {
			t.Log("unexpected result for case", i, c.lib, "\n\texpected:\t", c.bye, "\n\tgot:\t\t", bye)
			t.Fail()
		}
		if len(players) != c.lenAfter {
			t.Log("unexpected len(players) for case", i, c.lib, "\n\texpected:\t", c.lenAfter, "\n\tgot:\t\t", len(players))
			t.Fail()
		}

	}
}

func makeGameResult(table, p1, p2 string, vp1, vp2 int) Game {
	return Game{
		TableID: table,
		Results: [2]Result{
			{PlayerID: p1, VictoryPoints: vp1},
			{PlayerID: p2, VictoryPoints: vp2},
		},
	}
}
