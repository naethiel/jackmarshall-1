package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"github.com/HouzuoGuo/tiedot/data"
	"github.com/chibimi/jackmarshall/auth"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/cors"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

func main() {

	cfg := Configuration{}
	err := envconfig.Process("app", &cfg)
	if err != nil {
		log.Fatalln("unable to read configuration from env:", err)
	}
	fmt.Println(cfg)

	db, err := mgo.Dial("127.0.0.1:27017")

	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	roles := []string{"organizer"}
	router := httprouter.New()
	router.GET("/api/tournaments", auth.NewAuthHandler(NewListTournamentHandler(db), roles, cfg.Secret))
	router.GET("/api/tournaments/:id", auth.NewAuthHandler(NewGetTournamentHandler(db), roles, cfg.Secret))
	router.POST("/api/tournaments", auth.NewAuthHandler(NewCreateTournamentHandler(db), roles, cfg.Secret))
	router.PUT("/api/tournaments/:id", auth.NewAuthHandler(NewUpdateTournamentHandler(db), roles, cfg.Secret))
	router.DELETE("/api/tournaments/:id", auth.NewAuthHandler(NewDeleteTournamentHandler(db), roles, cfg.Secret))

	router.GET("/api/tournaments/:id/round", auth.NewAuthHandler(NewCreateRoundHandler(db), roles, cfg.Secret))
	router.GET("/api/tournaments/:id/results", auth.NewAuthHandler(NewGetResultsHandler(db), roles, cfg.Secret))

	// Initialize the middleware stack
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"accept", "authorization", "content-type"},
	})
	stack := negroni.New()
	stack.Use(cors)
	//	stack.Use(negroni.NewLogger())
	stack.Use(negroni.NewRecovery())
	stack.UseHandler(router)

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), stack))
}

func createTestTournament(database *data.Collection, name string, nbPlayers int, nbTables int, nbScenario int) {
	factions := []string{"Cygnar", "Trollbloods", "Legion of Everblight", "Cryx", "Khador"}
	tournament := Tournament{}
	players := []Player{}
	tables := []Table{}

	tournament.Name = name
	tournament.Rounds = []Round{}

	nbFaction := len(factions)
	lists := []List{List{Caster: "caster1"}, List{Caster: "caster2"}}
	for i := 0; i < nbPlayers; i++ {
		players = append(players, Player{
			Name:    "player" + fmt.Sprintf("%d", i),
			Faction: factions[rand.Intn(nbFaction)],
			Lists:   lists,
		})
	}

	for i := 0; i < nbTables; i++ {
		tables = append(tables, Table{Name: "table" + fmt.Sprintf("%d", i), Scenario: "scenario" + fmt.Sprintf("%d", i%nbScenario)})
	}

	tournament.Tables = tables
	tournament.Players = players

	data, err := json.Marshal(tournament)
	if err != nil {
		fmt.Println("ERROR MARSHALL")
	}

	_, err = database.Insert(data)
	if err != nil {
		fmt.Println("ERROR INSERT")

	}
}

func testAssignement() {
	tournament := Tournament{}
	players := []Player{}
	tables := []Table{}

	nbPlayers := 8
	nbTables := 4
	nbScenario := 4
	nbRounds := 2

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

		createRound(pairings, tables, &round)

		for _, g := range round.Games {
			g.Results[rand.Intn(2)].VictoryPoints = 1
		}

		tournament.Rounds = append(tournament.Rounds, round)
	}
	displayTournament(tournament)
	fmt.Println("FINI ! ")
}

func displayTournament(tournament Tournament) {
	for _, round := range tournament.Rounds {
		fmt.Println(round.String())
	}
}
