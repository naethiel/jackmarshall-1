package main

import (
	"log"
 	"net/http"

	"github.com/HouzuoGuo/tiedot/db"

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
	databasePath := "/database/jackmarshall"
	database, err := db.OpenDB(databasePath)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	database.Create("Tournaments")

	router := httprouter.New()
	router.GET("/tournaments", NewListTournamentHandler(database))
	router.POST("/tournaments", NewCreateTournamentHandler(database))
	router.NotFound = http.FileServer(http.Dir("app"))

	// Initialize the middleware stack
	stack := negroni.New()
	stack.Use(negroni.NewLogger())
	stack.Use(negroni.NewRecovery())
	stack.Use(negroni.HandlerFunc(ContentTypeHandler))
	stack.UseHandler(router)

	log.Fatalln(http.ListenAndServe(":8080", stack))
}
