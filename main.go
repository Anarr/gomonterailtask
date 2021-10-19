package main

import (
	"github.com/Anarr/gomonterailtask/api/api"
	"github.com/Anarr/gomonterailtask/config"
	"log"
)

func main() {
	//load configuration
	config, err := config.Init("development")

	if err != nil {
		log.Fatal(err)
	}

	api.Run(config)
}