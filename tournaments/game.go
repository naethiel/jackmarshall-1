package main

type Game struct {
	Table   Table     `json:"table" bson:"table"`
	Results [2]Result `json:"results" bson:"results"`
}
