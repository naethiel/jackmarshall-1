package main

type Result struct {
	Player            Player `json:"player"`
	List              string `json:"list"`
	VictoryPoints     int    `json:"victory_points"`
	ScenarioPoints    int    `json:"scenario_points,string"`
	DestructionPoints int    `json:"destruction_points,string"`
	SoS               int    `json:"sos"`
	Bye               bool   `json:"bye"`
}
