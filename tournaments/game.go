package tournaments

import "fmt"

type Game struct {
	Table   Table     `json:"table" bson:"table"`
	Results [2]Result `json:"results" bson:"results"`
}

func (g Game) String() string {
	return fmt.Sprintf("Table %s: %s (%d) vs %s (%d)", g.Table.Name, g.Results[0].Player.Name, g.Results[0].VictoryPoints, g.Results[1].Player.Name, g.Results[1].VictoryPoints)
}
