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
		logger.Log("level", "error", "error", err)
	}
	defer db.Close()

	roles := []string{"organizer"}
	router := httprouter.New()
	router.GET("/tournaments", auth.NewAuthHandler(NewListTournamentHandler(db), roles, cfg.Secret))
	router.GET("/tournaments/:id", auth.NewAuthHandler(NewGetTournamentHandler(db), roles, cfg.Secret))
	router.POST("/tournaments", auth.NewAuthHandler(NewCreateTournamentHandler(db), roles, cfg.Secret))
	router.PUT("/tournaments/:id", auth.NewAuthHandler(NewUpdateTournamentHandler(db), roles, cfg.Secret))
	router.DELETE("/tournaments/:id", auth.NewAuthHandler(NewDeleteTournamentHandler(db), roles, cfg.Secret))
	router.GET("/tournaments/:id/round", auth.NewAuthHandler(NewCreateRoundHandler(db), roles, cfg.Secret))
	router.GET("/tournaments/:id/results", auth.NewAuthHandler(NewGetResultsHandler(db), roles, cfg.Secret))

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
	stack.UseHandler(router)

	logger.Log("level", "info", "msg", "Server running", "port", cfg.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), stack)
	if err != nil {
		logger.Log("level", "error", "error", "err")
	}
}
