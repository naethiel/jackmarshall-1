package main

import (
	"fmt"
	"net/http"
	"os"

	mgo "gopkg.in/mgo.v2"

	"github.com/chibimi/jackmarshall/auth"
	"github.com/go-kit/kit/log"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/cors"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := log.With(log.NewLogfmtLogger(os.Stdout), "ts", log.DefaultTimestamp, "caller", log.DefaultCaller)

	cfg := Configuration{}
	err := envconfig.Process("app", &cfg)
	if err != nil {
		logger.Log("level", "error", "msg", "unable to read configuration from env", "error", err)
	}

	db, err := mgo.Dial(fmt.Sprintf("%s:%d", cfg.MongoAddr, cfg.MongoPort))
	if err != nil {
		logger.Log("level", "error", "msg", "unable to connect to database", "error", err)
	}
	defer db.Close()

	router := httprouter.New()
	router.GET("/tournaments", NewListTournamentHandler(db, logger))
	router.GET("/tournaments/:id", NewGetTournamentHandler(db, logger))
	router.POST("/tournaments", NewCreateTournamentHandler(db, logger))
	router.PUT("/tournaments/:id", NewUpdateTournamentHandler(db, logger))
	router.DELETE("/tournaments/:id", NewDeleteTournamentHandler(db, logger))
	router.GET("/tournaments/:id/round", NewCreateRoundHandler(db, logger))
	router.GET("/tournaments/:id/results", NewGetResultsHandler(db, logger))

	// Initialize the middleware stack
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"accept", "authorization", "content-type"},
	})
	stack := negroni.New()
	stack.Use(cors)
	stack.Use(negroni.NewLogger())
	stack.Use(negroni.NewRecovery())
	stack.Use(auth.NewAuthMiddleware(cfg.Secret))
	stack.UseHandler(router)

	// collection := db.DB("jackmarshall").C("tournament")
	// tournament := tournaments.NewTestTournament(8, 4, 4, true)
	// tournament.Owner = 1
	// tournament.ID = bson.NewObjectId()
	// collection.Insert(&tournament)

	logger.Log("level", "info", "msg", "Server running", "port", cfg.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), stack)
	if err != nil {
		logger.Log("level", "error", "error", "err")
	}

}
