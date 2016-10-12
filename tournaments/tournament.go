package main

import (
	"time"
)

type Tournament struct {
	Id        string    `json:"id" update:"nonzero"`
	Name      string    `json:"name" create:"nonzero" update:"nonzero"`
	Format    int       `json:"format" create:"nonzero" update:"nonzero"`
	Slots     int       `json:"slots" create:"nonzero" update:"nonzero"`
	FeeAmount float64   `json:"fee_amount"`
	Date      time.Time `json:"date" create:"nonzero" update:"nonzero"`
	Players   []Player  `json:"players"`
	Tables    []Table   `json:"tables"`
	Rounds    []Round   `json:"rounds" create:"max=0"`
}

func NewTournament() *Tournament {
	t := new(Tournament)
	t.Players = []Player{}
	t.Tables = []Table{}
	t.Rounds = []Round{}
	return t
}
