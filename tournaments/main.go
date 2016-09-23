package main

import (
	"log"
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
	router.NotFound = http.FileServer(http.Dir("front"))

	// Initialize the middleware stack
	stack := negroni.New()
	stack.Use(negroni.NewLogger())
	stack.Use(negroni.NewRecovery())
	stack.UseHandler(router)

	log.Fatalln(http.ListenAndServe(":8080", stack))
}
