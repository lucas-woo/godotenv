package main

import (
	"log"
	"github.com/lucas-woo/godotenv"
)
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
}