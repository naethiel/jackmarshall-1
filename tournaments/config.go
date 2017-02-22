package main

type Configuration struct {
	PasswordCost int    `envconfig:"password_cost"`
	Port         int    `envconfig:"port"`
	MongoAddr    string `envconfig:"DATABASE_PORT_27017_TCP_ADDR"`
	MongoPort    int    `envconfig:"DATABASE_PORT_27017_TCP_PORT"`
	Secret       string `envconfig:"secret"`
}
