package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chibimi/jackmarshall/auth"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/cors"
	"menteslibres.net/gosexy/redis"
)

func main() {
	configuration := Configuration{}
	err := envconfig.Process("app", &configuration)
	if err != nil {
		log.Fatalln("unable to read configuration from env:", err)
	}

	database := redis.New()
	err = database.Connect(configuration.RedisAddr, uint(configuration.RedisPort))
	if err != nil {
		log.Fatalln("unable to connect to redis database:", err)
	}

	router := httprouter.New()
	router.GET("/users", NewUserListHandler(database))
	router.POST("/users", NewUserCreateHandler(database, configuration))
	router.POST("/organizer", NewOrganizerCreateHandler(database, configuration))
	router.GET("/users/:id", NewUserShowHandler(database))
	router.PUT("/users/:id", NewUserUpdateHandler(database, configuration))
	router.POST("/login", NewUserLoginHandler(database, configuration))
	router.POST("/refresh", NewRefreshTokenHandler(database, configuration))
	router.DELETE("/refresh", NewInvalidateRefreshTokenHandler(database, configuration))

	router.GET("/usersAuth", auth.NewAuthHandler(NewUserListHandler(database), []string{"roleTest"}, configuration.Secret))

	// Initialize the middleware stack
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"accept", "authorization", "content-type"},
	})
	stack := negroni.New()
	stack.Use(cors)
	stack.UseHandler(router)

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", configuration.Port), stack))
}
