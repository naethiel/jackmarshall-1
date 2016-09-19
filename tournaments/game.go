package main

type Game struct {
	Table   string   `json:"table" bson:"table"`
	Pairing []Result `json:"pairing" bson:"pairing"`
}
