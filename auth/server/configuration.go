package main

type Configuration struct {
	PasswordCost int    `envconfig:"password_cost"`
	Port         int    `envconfig:"port"`
	RedisAddr    string `envconfig:"redis_port_6379_tcp_addr"`
	RedisPort    int    `envconfig:"redis_port_6379_tcp_port"`
	Secret       string `envconfig:"secret"`
}
