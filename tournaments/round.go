package main

import "fmt"

type Round struct {
	Number int    `json:"number"`
	Games  []Game `json:"games"`
}

func getAvailableTables(pairs []Pair, tables []Table) map[Table]map[Pair]struct{} {
	var availableTables = make(map[Table]map[Pair]struct{})
	for _, table := range tables {
		availableTables[table] = make(map[Pair]struct{})
		for _, pair := range pairs {
			if pair.PlayedOn(table.Scenario) {
				continue
			}
			availableTables[table][pair] = struct{}{}
		}
	}
	return availableTables
}

func getAvailablePairs(pairs []Pair, tables []Table) map[Pair]map[Table]struct{} {
	var availablePairs = make(map[Pair]map[Table]struct{})
	for _, pair := range pairs {
		availablePairs[pair] = make(map[Table]struct{})
		for _, table := range tables {
			if pair.PlayedOn(table.Scenario) {
				continue
			}
			availablePairs[pair][table] = struct{}{}
		}
	}
	return availablePairs
}

func CreateRound(pairs []Pair, tables []Table, round *Round) {

	var availableTables = getAvailableTables(pairs, tables)
	var availablePairs = getAvailablePairs(pairs, tables)

	games := round.Games

Selection:
	for len(availablePairs) != 0 {

		// Find tables with 1 choice available.
		for table, pairs := range availableTables {

			if len(pairs) == 1 {

				var pair Pair
				for pair = range pairs {
					break
				}
				fmt.Println("PAIR ")
				fmt.Println(pair)
				games = append(games, Game{
					Table: table,
					Results: [2]Result{
						Result{
							Player: *pair[0],
						},
						Result{
							Player: *pair[1],
						},
					},
				})

				// Delete the table from the available tables and pairs sets.
				delete(availableTables, table)
				for pair, tables := range availablePairs {
					delete(tables, table)
					availablePairs[pair] = tables
				}
				delete(availablePairs, pair)
				for table, pairs := range availableTables {
					delete(pairs, pair)
					availableTables[table] = pairs
				}

				continue Selection
			}
		}

		// Find pairs with 1 choice available.
		for pair, tables := range availablePairs {

			if len(tables) == 1 {

				var table Table
				for table = range tables {
					break
				}
				games = append(games, Game{
					Table: table,
					Results: [2]Result{
						Result{
							Player: *pair[0],
						},
						Result{
							Player: *pair[1],
						},
					},
				})

				// Delete the table from the available tables and pairs sets.
				delete(availableTables, table)
				for pair, tables := range availablePairs {
					delete(tables, table)
					availablePairs[pair] = tables
				}
				delete(availablePairs, pair)
				for table, pairs := range availableTables {
					delete(pairs, pair)
					availableTables[table] = pairs
				}

				continue Selection
			}
		}
		break
	}
	//BRUTE FORCE !
	remainingTables := getTablesKeys(availableTables)
	remainingPairs := getPairsKeys(availablePairs)
	min := 1
	minScore := &min
	fmt.Printf("remaningTable : %d\n", len(remainingTables))
	solution := assignTables(make([]Game, len(remainingTables)), 0, remainingTables, len(remainingTables), remainingPairs, make([]Game, len(remainingTables)), minScore)
	games = append(games, solution...)

	round.Games = games
	return
}

func assignTables(attempt []Game, position int, remainingTables []Table, nbTables int, pairings []Pair, solution []Game, minScore *int) []Game {
	if len(remainingTables) == 0 {
		*minScore = CalculateScore(attempt)
		copy(solution, attempt)
		if *minScore == 0 {
			fmt.Println("SOLUTION PARFAITE")
			return solution
		}
	}

	for i, table := range remainingTables {
		attempt[position] = Game{
			Table: table,
			Results: [2]Result{
				Result{
					Player: *pairings[position][0],
				},
				Result{
					Player: *pairings[position][1],
				},
			},
		}
		// fmt.Println("ATTEMPT")
		// fmt.Println(attempt)
		if CalculateScore(attempt) >= *minScore {
			attempt = append(attempt[:position], make([]Game, nbTables-position)...)
		} else {
			remainingTablesCopy := make([]Table, len(remainingTables))
			copy(remainingTablesCopy, remainingTables)

			remainingTablesCopy = append(remainingTablesCopy[:i], remainingTablesCopy[i+1:]...)

			assignTables(attempt, position+1, remainingTablesCopy, nbTables, pairings, solution, minScore)
		}
	}
	//fmt.Println("SOLUTION DE MERDE")
	return solution
}

func CalculateScore(attempt []Game) (res int) {
	res = 0
	for _, game := range attempt {
		for _, result := range game.Results {
			for _, playerGame := range result.Player.Games {
				if game.Table == playerGame.Table {
					res += 50
				} else if game.Table.Scenario == playerGame.Table.Scenario {
					res += 10
				}
			}
		}
	}
	return
}

func contains(s []string, e string) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}

func getTablesKeys(availableTables map[Table]map[Pair]struct{}) []Table {
	keys := make([]Table, 0, len(availableTables))
	for k := range availableTables {
		keys = append(keys, k)
	}
	return keys
}

func getPairsKeys(availablePairs map[Pair]map[Table]struct{}) []Pair {
	keys := make([]Pair, 0, len(availablePairs))
	for k := range availablePairs {
		keys = append(keys, k)
	}
	return keys
}
