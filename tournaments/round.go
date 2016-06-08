package main

type Round struct {
	Number int    `json:"number"`
	Games  []Game `json:"games"`
}

func CreateRound(pairs []Pair, tables []Table, round *Round) {

	var availableTables = make(map[Table]map[Pair]struct{})
	for _, table := range tables {
		availableTables[table] = make(map[Pair]struct{})
		for _, pair := range pairs {
			if pair.PlayedOn(table) {
				continue
			}
			availableTables[table][pair] = struct{}{}
		}
	}

	var availablePairs = make(map[Pair]map[Table]struct{})
	for _, pair := range pairs {
		availablePairs[pair] = make(map[Table]struct{})
		for _, table := range tables {
			if pair.PlayedOn(table) {
				continue
			}
			availablePairs[pair][table] = struct{}{}
		}
	}
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
							Player: pair[0].Name,
						},
						Result{
							Player: pair[1].Name,
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
							Player: pair[0].Name,
						},
						Result{
							Player: pair[1].Name,
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

		// Assign the first pair with at least 2 tables available to the first
		// available table.
		for pair, tables := range availablePairs {

			if len(tables) > 1 {

				var table Table
				for table = range tables {
					break
				}

				games = append(games, Game{
					Table: table,
					Results: [2]Result{
						Result{
							Player: pair[0].Name,
						},
						Result{
							Player: pair[1].Name,
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

		// There is only suckers left, assign them to the first remaining table.
		for pair := range availablePairs {

			var table Table
			for table = range availableTables {
				break
			}

			games = append(games, Game{
				Table: table,
				Results: [2]Result{
					Result{
						Player: pair[0].Name,
					},
					Result{
						Player: pair[1].Name,
					},
				},
			})

			// Delete the table from the available tables.
			delete(availableTables, table)
		}

		break
	}
	round.Games = games
	return
}
