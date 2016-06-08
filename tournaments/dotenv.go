// +build !docker

package main

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("jackmarshall.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
