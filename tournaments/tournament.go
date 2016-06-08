package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Tournament struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty" update:"nonzero"`
	Owner     int           `json:"owner" bson:"owner"`
	Name      string        `json:"name" create:"nonzero" update:"nonzero"`
	Format    int           `json:"format" create:"nonzero" update:"nonzero"`
	Slots     int           `json:"slots" create:"nonzero" update:"nonzero"`
	FeeAmount float32       `json:"fee_amount"`
	Date      time.Time     `json:"date" create:"nonzero" update:"nonzero"`
	Players   []Player      `json:"players" create:"max=0"`
	Tables    []Table       `json:"tables" create:"max=0"`
	Rounds    []Round       `json:"rounds" create:"max=0"`
}
