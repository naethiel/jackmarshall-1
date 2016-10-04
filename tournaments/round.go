package main

import "fmt"

type Round struct {
	Number int    `json:"number"`
	Games  []Game `json:"games"`
}

func (r *Round) String() (s string) {
	s += fmt.Sprintf("ROUND %d : %d\n", r.Number, GetFitness(r.Games))
	for i, game := range r.Games {
		s += fmt.Sprintf("GAME %d :\t%s  %s\t", i, game.Table.Name, game.Table.Scenario)
		s += fmt.Sprintf("%s vs %s\n", game.Results[0].Player.Name, game.Results[1].Player.Name)
	}
	return
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

func createRound(pairs []Pair, tables []Table, round *Round) {

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

	remainingTables := getTablesKeys(availableTables)
	remainingPairs := getPairsKeys(availablePairs)

	if len(remainingPairs) == 1 {
		games = append(games, Game{
			Table: remainingTables[0],
			Results: [2]Result{
				Result{
					Player: *remainingPairs[0][0],
				},
				Result{
					Player: *remainingPairs[0][1],
				},
			},
		})

	} else {
		solution := GetBest(remainingTables, remainingPairs, 10)
		games = append(games, solution...)
	}

	round.Games = games
	return
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
