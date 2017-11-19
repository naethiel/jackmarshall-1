package tournaments

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

var debug = false

type Genetic struct {
	PopSize int
	NbBest  int
	Delay   time.Duration
}

type Population []*Individual

type Individual struct {
	Fitness int
	Players []*Player
}

func (in *Individual) CalcFitness() {
	fit := 0

	for i := 0; i < len(in.Players); i++ {
		if i == len(in.Players)-1 && in.Players[i].HadBye() {
			if debug {
				fmt.Printf("%s has already had a bye\n", in.Players[i].Name)
			}
			fit += 10000
			break
		}

		if in.Players[i].PlayedAgainst(in.Players[i+1]) {
			if debug {
				fmt.Printf("%s has already played vs %s\n", in.Players[i].Name, in.Players[i+1].Name)
			}
			fit += 10000
		}

		diff := math.Abs(float64(in.Players[i].VictoryPoints() - in.Players[i+1].VictoryPoints()))
		if diff > 1 {
			if debug {
				fmt.Printf("too many point between %s and %s\n", in.Players[i].Name, in.Players[i+1].Name)
			}
			fit += 10000
			// } else if diff == 1 && in.players[i].HadSousApp(){
			// 	if debug {
			// 		fmt.Printf("sous app entre %s et %s\n", in.players[i].Name, in.players[i+1].Name)
			// 	}
			// 	fit += 100
		} else if diff == 1 {
			if debug {
				fmt.Printf("sous app entre %s et %s\n", in.Players[i].Name, in.Players[i+1].Name)
			}
			fit += 100
		}

		// Origin
		if in.Players[i].Origin == in.Players[i+1].Origin {
			// if debug {
			// 	fmt.Printf("origine entre %s et %s\n", players[i].Name, players[i+1].Name)
			// }
			fit += 1
		}
		i++
	}
	in.Fitness = fit
}

func (in *Individual) Mutate() {
	i := rand.Intn(len(in.Players) - 1)
	temp := in.Players[i+1]
	in.Players[i+1] = in.Players[i]
	in.Players[i] = temp
}

func (in *Individual) String() string {
	s := fmt.Sprintf("%d ==>", in.Fitness)
	for _, p := range in.Players {
		s = fmt.Sprintf("%s%s, ", s, p.Name)
	}
	return s
}

func (p Population) SortByFitness() {
	sort.Slice(p, func(i int, j int) bool {
		return p[i].Fitness < p[j].Fitness
	})
}

func (p Population) CalcFitness() {
	for _, in := range p {
		in.CalcFitness()
	}
}

func (p Population) Mutate() {
	for _, in := range p {
		in.Mutate()
	}
}

func (p Population) String() string {
	s := ""
	for _, in := range p {
		s = fmt.Sprintf("%s\n%s", s, in.String())
	}
	return s
}

func NextGeneration(pop Population, nbBest int) Population {
	next := Population{}
	for i := 0; i < nbBest; i++ {
		next = append(next, pop[i])
	}
	nbLucky := len(pop) - nbBest
	for i := 0; i < nbLucky; i++ {
		next = append(next, pop[rand.Intn(nbLucky)])
	}
	return next
}

func NewInitialPopulation(src []*Player, size int) Population {
	pop := Population{}
	for i := 0; i < size; i++ {
		pop = append(pop, generateIndividual(src))
	}
	return pop
}

func generateIndividual(p []*Player) *Individual {
	players := make([]*Player, len(p))
	copy(players, p)
	for i := range players {
		var j = rand.Intn(len(players) - 1)
		players[i], players[j] = players[j], players[i]
	}
	//Sort player list by victory points
	sort.Slice(players, func(i int, j int) bool {
		return players[i].VictoryPoints() > players[j].VictoryPoints()
	})
	return &Individual{
		Fitness: math.MaxInt32,
		Players: players,
	}
}

func (s *Genetic) Solve(src []*Player) Individual {
	rand.Seed(time.Now().UnixNano())

	pop := NewInitialPopulation(src, s.PopSize)
	best := Individual{
		Fitness: math.MaxInt32,
	}
	timeout := false
	go func() {
		time.Sleep(s.Delay)
		timeout = true
	}()

	for !timeout && best.Fitness != 0 {
		pop.CalcFitness()

		pop.SortByFitness()
		if pop[0].Fitness < best.Fitness {
			fmt.Println("New best found ", pop[0].Fitness)
			best = *pop[0]
		}
		pop = NextGeneration(pop, s.NbBest)
		pop.Mutate()
	}
	return best
}
