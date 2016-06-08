package main

import "gopkg.in/mgo.v2/bson"

type Table struct {
	Name     string        `json:"name"`
	Scenario bson.ObjectId `json:"scenario" bson:"scenario"`
}
