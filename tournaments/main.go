package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/HouzuoGuo/tiedot/data"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

func main() {
	databasePath := "database"
	database, err := data.OpenCollection(databasePath)
	if err != nil {
		panic(err)
	}

	router := httprouter.New()
	router.GET("/api/tournaments", NewListTournamentHandler(database))
	router.GET("/api/tournaments/:id", NewGetTournamentHandler(database))
	router.POST("/api/tournaments", NewCreateTournamentHandler(database))
	router.PUT("/api/tournaments/:id", NewUpdateTournamentHandler(database))
	router.DELETE("/api/tournaments/:id", NewDeleteTournamentHandler(database))

	router.GET("/api/tournaments/:id/round", NewCreateRoundHandler(database))

	router.NotFound = http.FileServer(http.Dir("front"))

	// Initialize the middleware stack
	stack := negroni.New()
	stack.Use(negroni.NewLogger())
	stack.Use(negroni.NewRecovery())
	stack.UseHandler(router)

	testAssignement()

	log.Fatalln(http.ListenAndServe(":8080", stack))
}

func testAssignement() {
	tournament := Tournament{}
	players := []Player{}
	tables := []Table{}

	nbPlayers := 32
	nbTables := 16
	nbScenario := 8
	nbRounds := 5

	for i := 0; i < nbPlayers; i++ {
		players = append(players, Player{Name: "player" + fmt.Sprintf("%d", i)})
	}

	for i := 0; i < nbTables; i++ {
		tables = append(tables, Table{Name: "table" + fmt.Sprintf("%d", i), Scenario: "scenario" + fmt.Sprintf("%d", i%nbScenario)})
	}

	tournament.Tables = tables
	tournament.Players = players

	for i := 0; i < nbRounds; i++ {
		round := Round{
			Number: i,
			Games:  []Game{},
		}
		var pairings = CreatePairs(players, tournament, &round)

		CreateRound(pairings, tables, &round)
		fmt.Printf("ROUND %d : %d\n", round.Number, CalculateScore(round.Games))

		for _, g := range round.Games {
			g.Results[rand.Intn(2)].VictoryPoints = 1

		}
		// for _, game := range round.Games {
		// 	//fmt.Println(game.Table, game.Results[0], game.Results[1])
		// }

		tournament.Rounds = append(tournament.Rounds, round)
		//fmt.Printf("%+v", tournament)
	}
	fmt.Println("FINI ! ")
}

func calculateRoundScore(round Round, tournament Tournament) {
	score := 0
	for _, game := range round.Games {
		for _, player := range tournament.Players {
			if player.Name == game.Results[0].Player.Name || player.Name == game.Results[1].Player.Name {
				for _, playerGame := range player.Games {
					if playerGame.Table.Scenario == game.Table.Scenario {
						score += 10
					}
				}
			}
		}
	}

	fmt.Printf("ROUND %d : %d\n", round.Number, score)
}
