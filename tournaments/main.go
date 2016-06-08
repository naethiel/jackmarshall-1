package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"git.elwinar.com/jackmarshall/auth"
	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/cors"
	"gopkg.in/mgo.v2"
)

//Set the content Type of the response to Json
func ContentTypeHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	rw := negroni.NewResponseWriter(w)
	rw.Before(func(rw negroni.ResponseWriter) {
		if rw.Status() == http.StatusOK {
			rw.Header().Set("Content-Type", "application/json")
		}
	})
	next(rw, r)
}

func main() {
	configuration := Configuration{}
	err := envconfig.Process("app", &configuration)
	if err != nil {
		log.Fatalln("unable to read configuration from env:", err)
	}

	database, err := mgo.Dial(configuration.MongoAddr + ":" + strconv.Itoa(configuration.MongoPort))
	if err != nil {
		log.Fatalln(err)
	}
	defer database.Close()

	router := httprouter.New()
	router.GET("/scenario", NewListScenarioHandler(database))
	router.GET("/scenario/:id", NewGetScenarioHandler(database))
	router.POST("/scenario", NewCreateScenarioHandler(database))
	router.PUT("/scenario/:id", NewUpdateScenarioHandler(database))
	router.DELETE("/scenario/:id", NewDeleteScenarioHandler(database))

	router.GET("/tables", NewListTableHandler(database))
	router.GET("/players", NewListPlayerHandler(database))

	roles := []string{"roleTest"}
	router.GET("/playersAuth", auth.NewAuthHandler(NewListPlayerHandler(database), roles, configuration.Secret))
	router.GET("/tablesAuth", auth.NewAuthHandler(NewListTableHandler(database), roles, configuration.Secret))
	router.GET("/tournamentsAuth", auth.NewAuthHandler(NewListTournamentHandler(database), roles, configuration.Secret))
	router.POST("/tournamentsAuth", auth.NewAuthHandler(NewCreateTournamentHandler(database), roles, configuration.Secret))

	router.GET("/tournaments/:id/round", NewCreateRoundHandler(database))

	router.GET("/tournaments", NewListTournamentHandler(database))
	router.GET("/tournaments/:id", NewGetTournamentHandler(database))
	router.POST("/tournaments", NewCreateTournamentHandler(database))
	router.PUT("/tournaments/:id", NewUpdateTournamentHandler(database))
	router.DELETE("/tournaments/:id", NewDeleteTournamentHandler(database))

	// Initialize the middleware stack
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8000"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"accept", "authorization", "content-type"},
	})
	stack := negroni.New()
	stack.Use(cors)
	stack.Use(negroni.NewLogger())
	stack.Use(negroni.NewRecovery())
	stack.Use(negroni.HandlerFunc(ContentTypeHandler))
	stack.UseHandler(router)

	//testAssignement()

	log.Fatalln(http.ListenAndServe(":8080", stack))
}

//TODO test, to be deleted
func testAssignement() {
	tournament := Tournament{}
	players := []Player{}
	tables := []Table{}

	nbPlayers := 15
	nbTables := 8
	nbRounds := 3

	for i := 0; i < nbPlayers; i++ {
		players = append(players, Player{Name: "player" + fmt.Sprintf("%d", i+1)})
	}

	for i := 0; i < nbTables; i++ {
		tables = append(tables, Table{Name: "table" + fmt.Sprintf("%d", i+1)})
	}

	tournament.Tables = tables
	tournament.Players = players

	for i := 0; i < nbRounds; i++ {
		fmt.Println("round", i+1)
		fmt.Println("")
		round := Round{
			Games: []Game{},
		}
		var pairings = CreatePairs(players, tournament, &round)

		CreateRound(pairings, tables, &round)

		for g := range round.Games {
			round.Games[g].Results[rand.Intn(2)].VictoryPoints = 1
		}
		for _, game := range round.Games {
			fmt.Println(game.Table, game.Results[0], game.Results[1])
		}
		tournament.Rounds = append(tournament.Rounds, round)
		fmt.Println(tournament)
	}
}
