package solver

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var debug = false

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Mutable interface {
	CalcFitness() float64
	Mutate(float64) *Individual
	NewIndividual() *Individual
	String() string
}

type Solver struct {
	PopulationSize   int
	MaxIterations    int
	NumberOfChildren int
	RandomSwapRate   float64
}

type Population []*Individual

type Individual struct {
	Fitness float64
	Genes   Mutable
}

func (in *Individual) String() string {
	return fmt.Sprintf("%f ==> %s\n", in.Fitness, in.Genes.String())
}

func NewInitialPopulation(src Mutable, size int) Population {
	pop := Population{}
	for i := 0; i < size; i++ {
		pop = append(pop, src.NewIndividual())
	}
	return pop
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

func (p Population) SortByFitness() {
	sort.Slice(p, func(i int, j int) bool {
		return p[i].Fitness < p[j].Fitness
	})
}

func (p Population) CalcFitness() {
	for _, in := range p {
		in.Fitness = in.Genes.CalcFitness()
	}
}

func (p Population) Mutate(randomSwapRate float64) Population {
	// for _, in := range p {
	// 	in.Genes.Mutate(randomSwapRate)
	// }
	var child = make(Population, len(p))
	copy(child, p)

	for i := 0; i < len(child); i++ {
		child[i] = p[i].Genes.Mutate(randomSwapRate)
	}
	return child
}

func (s *Solver) Solve(src Mutable) (Individual, int) {
	pop := NewInitialPopulation(src, s.PopulationSize)
	pop.CalcFitness()
	n := 0
	for ; n < s.MaxIterations && pop[0].Fitness != 0; n++ {
		pop = append(pop, pop.Mutate(s.RandomSwapRate)...)
		pop.CalcFitness()
		sort.Slice(pop, func(i, j int) bool {
			return pop[i].Fitness < pop[j].Fitness
		})

		pop = pop[:s.PopulationSize]
		// fmt.Println(s, fitness(population[0]), population[0])
	}

	return *pop[0], n
}
