package main

//Result represents the result of a player for a game
type Result struct {
	Player            Player `json:"player"`
	List              int    `json:"list"`
	VictoryPoints     int    `json:"victory_points"`
	ScenarioPoints    int    `json:"scenario_points"`
	DestructionPoints int    `json:"destruction_points"`
	Bye               bool   `json:"bye"`
}
