package main

import (
	"gopkg.in/mgo.v2/bson"
)

type Scenario struct {
	ID   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name string        `json:"name" validate:"nonzero"`
	Year int           `json:"year"`
	Link string        `json:"link"`
}
