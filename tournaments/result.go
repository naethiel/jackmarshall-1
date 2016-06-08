package main

type Result struct {
	Player            string `json:"player" bson:"player"`
	List              int    `json:"list"`
	VictoryPoints     int    `json:"victory_points" bson:"victory_points"`
	ScenarioPoints    int    `json:"scenario_points" bson:"scenario_points"`
	DestructionPoints int    `json:"destruction_points" bson:"destruction_points"`
	Buy               bool   `json:"buy" bson: "buy"`
}
