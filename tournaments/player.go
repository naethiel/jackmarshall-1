package main

type Player struct {
	Name     string   `json:"name"`
	Faction  string   `json:"faction"`
	PayedFee bool     `json:"payed_fee"`
	Lists    []string `json:"lists"`
	Leave    bool     `json:"leave"`
}
