package main

import (
	"log"

	"github.com/htnk128/go-ddd-sample/pkg/address/external/rest"
)

func main() {
	s := rest.NewHttpServer()
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to listen and serve: %+v", err)
	}
}
