package main

import (
	"log"

	"github.com/azimov-mg/gometr/api"
)

func main() {
	log.Println("Starting GoMetr...")

	api.StartServer()

	log.Println("Shutting down GoMetr...")
}
