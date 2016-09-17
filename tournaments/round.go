package main

type Round struct {
	Number int    `json:"number"`
	Games  []Game `json:"games"`
}

func NewRoundFromPairing(pairings []Pairing, tables []string) (round Round) {
	round.Number = 0
	for i, v := range tables {
		g := Game{}
		r1 := Result{}
		r2 := Result{}

		r1.Player = pairings[i].Player1
		r2.Player = pairings[i].Player2

		g.Table = v
		g.Pairing = append(g.Pairing, r1)
		g.Pairing = append(g.Pairing, r2)
		round.Games = append(round.Games, g)
	}
	return
}
