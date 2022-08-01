package main

import (
	"log"

	"github.com/gaponovalexey/go-restapi/pkg/app/apiserver"
)

func main() {

	
	config := apiserver.NewConfig()
	s := apiserver.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
