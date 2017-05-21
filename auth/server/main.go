package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/go-kit/kit/log"
	"github.com/julienschmidt/httprouter"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/cors"
	"menteslibres.net/gosexy/redis"
)

func main() {
	logger := log.With(log.NewLogfmtLogger(os.Stdout), "ts", log.DefaultTimestamp, "caller", log.DefaultCaller)

	cfg := Configuration{}
	err := envconfig.Process("app", &cfg)
	if err != nil {
		logger.Log("level", "error", "msg", "unable to read configuration from env", "error", err)
	}

	database := redis.New()
	err = database.Connect(cfg.RedisAddr, cfg.RedisPort)
	if err != nil {
		logger.Log("level", "error", "msg", "unable to connect to database", "error", err)
	}
	defer database.Close()

	router := httprouter.New()
	router.GET("/users", NewUserListHandler(database, logger))
	router.POST("/users", NewUserCreateHandler(database, logger, cfg))
	router.POST("/organizer", NewOrganizerCreateHandler(database, logger, cfg))
	router.GET("/users/:id", NewUserShowHandler(database, logger))
	router.PUT("/users/:id", NewUserUpdateHandler(database, logger, cfg))
	router.POST("/login", NewUserLoginHandler(database, logger, cfg))
	router.POST("/refresh", NewRefreshTokenHandler(database, logger, cfg))
	router.DELETE("/refresh", NewInvalidateRefreshTokenHandler(database, logger, cfg))

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

	logger.Log("level", "info", "msg", "Server running", "port", cfg.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), stack)
	if err != nil {
		logger.Log("level", "error", "error", "err")
	}
}
