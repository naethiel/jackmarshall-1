package main

import (
	"log"
	"net/http"

	"github.com/HouzuoGuo/tiedot/data"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

//Set the content Type of the response to Json
func ContentTypeHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	rw := negroni.NewResponseWriter(w)
	rw.Before(func(rw negroni.ResponseWriter) {
		if rw.Status() == http.StatusOK {
			//	rw.Header().Set("Content-Type", "application/json")
		}
	})
	next(rw, r)
}

func main() {
	databasePath := "database"
	database, err := data.OpenCollection(databasePath)
	// database, err := db.OpenDB(databasePath)
	if err != nil {
		panic(err)
	}
	// defer database.Close()

	// database.Create("Tournaments")

	router := httprouter.New()
	router.GET("/api/tournaments", NewListTournamentHandler(database))
	router.GET("/api/tournaments/:id", NewGetTournamentHandler(database))
	router.POST("/api/tournaments", NewCreateTournamentHandler(database))
	router.PUT("/api/tournaments/:id", NewUpdateTournamentHandler(database))
	router.DELETE("/api/tournaments/:id", NewDeleteTournamentHandler(database))
	router.NotFound = http.FileServer(http.Dir("app"))

	// Initialize the middleware stack
	stack := negroni.New()
	stack.Use(negroni.NewLogger())
	stack.Use(negroni.NewRecovery())
	stack.Use(negroni.HandlerFunc(ContentTypeHandler))
	stack.UseHandler(router)

	log.Fatalln(http.ListenAndServe(":8080", stack))
}
