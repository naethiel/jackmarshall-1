package tournaments

import (
	"fmt"
	"math/rand"
	"time"
)

func suffle(t []Player) {
	for i := 0; i < len(t); i++ {
		j := rand.Intn(len(t))
		t[i], t[j] = t[j], t[i]
	}
}

func printBracket(b map[int][]Player) {

	for k, v := range b {
		s := fmt.Sprintf("Bracket %d: ", k)
		for _, p := range v {
			s = fmt.Sprintf("%s, %s", s, p.ID)
		}
		fmt.Println(s)

	}
}

func getPlayerIndex(p Player, tab []Player) int {
	for i, v := range tab {
		if v.ID == p.ID {
			return i
		}
	}
	return -1
}

func NewTestTournament(nbPlayer, nbTable, nbScenario int, withOrigin bool) *Tournament {
	t := NewTournament()

	var origins = []string{"whag", "loin", "uchro", "ludo", "uchro", ""}
	var factions = []string{"cygnar", "cryx", "legion", "skorne", "trollbloods", ""}

	for i := 0; i < nbPlayer; i++ {
		o := ""
		if withOrigin {
			o = origins[rand.Intn(len(origins))]
		}
		p := Player{
			ID:      "player" + fmt.Sprintf("%d", i),
			Name:    "player" + fmt.Sprintf("%d", i),
			Origin:  o,
			Faction: factions[rand.Intn(len(factions))],
		}
		t.Players[p.ID] = p
	}

	for i := 0; i < nbTable; i++ {
		table := Table{
			ID:       "table" + fmt.Sprintf("%d", i),
			Name:     "table" + fmt.Sprintf("%d", i),
			Scenario: "scenario" + fmt.Sprintf("%d", i%nbScenario),
		}
		t.Tables[table.ID] = table
	}
	t.Date = time.Now().Add(time.Hour)

	return t
}
func (t *Tournament) CheckRound(index int) ([]string, []string, []string, []string) {
	round := t.Rounds[index]

	pairing := []string{}
	origin := []string{}
	sousApp := []string{}
	bye := []string{}

	verif := map[string]int{}
	for _, g := range round.Games {
		p0 := t.Players[g.Results[0].PlayerID]
		p1 := t.Players[g.Results[1].PlayerID]
		if p0.PlayedAgainst(p1.ID) {
			pairing = append(pairing, fmt.Sprintf("%s vs %s", p0.ID, p1.ID))
		}
		if p0.Origin != "" && p0.Origin == p1.Origin {
			origin = append(origin, fmt.Sprintf("%s vs %s (%s)", p0.ID, p1.ID, p0.Origin))
		}
		if g.Results[0].Bye && p0.HadBye() {
			bye = append(bye, p0.ID)
		}
		if g.Results[0].SousApp && p0.HadSousApp() {
			sousApp = append(sousApp, p0.ID)
		}
		if g.Results[1].SousApp && p1.HadSousApp() {
			sousApp = append(sousApp, p1.ID)
		}
		if _, ok := verif[g.Results[0].PlayerID]; ok {
			fmt.Println("doublon", round.Games)
		}
		if _, ok := verif[g.Results[1].PlayerID]; ok {
			fmt.Println("doublon", round.Games)
		}
		verif[g.Results[0].PlayerID] = 1
		verif[g.Results[1].PlayerID] = 1

	}
	//
	// fmt.Println("Pairing (", len(pairing), "):", pairing)
	// fmt.Println("Origin (", len(origin), "):", origin)
	// fmt.Println("Bye (", len(bye), "):", bye)
	// fmt.Println("SousApp (", len(sousApp), "):", sousApp)

	return pairing, origin, bye, sousApp

}
