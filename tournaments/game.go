package tournaments

import "fmt"

type Game struct {
	TableID string    `json:"table" bson:"table"`
	Results [2]Result `json:"results" bson:"results"`
}

func (g Game) String() string {
	return fmt.Sprintf("Table %s: %s (%d) vs %s (%d)", g.TableID, g.Results[0].PlayerID, g.Results[0].VictoryPoints, g.Results[1].PlayerID, g.Results[1].VictoryPoints)
}
