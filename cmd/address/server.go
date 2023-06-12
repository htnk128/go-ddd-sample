package main

import (
	"log"
	"net/http"
	"time"

	"github.com/htnk128/go-ddd-sample/pkg/address/external/rest"
)

const (
	addr         = ":8081"
	readTimeout  = time.Duration(5000)
	writeTimeout = time.Duration(5000)
)

func main() {
	router := rest.InitRouter()

	s := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  readTimeout * time.Second,
		WriteTimeout: writeTimeout * time.Second,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to listen and serve: %+v", err)
	}
}
