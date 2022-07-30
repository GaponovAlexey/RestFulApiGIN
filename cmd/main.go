package main

import (
	"fmt"
	"log"

	"github.com/gaponovalexey/go-restapi/internal/app/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
