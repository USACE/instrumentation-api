package main

import (
	"log"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/handler"
	"github.com/USACE/instrumentation-api/api/internal/server"
)

func main() {
	cfg := config.NewApiConfig()

	h := handler.NewApi(cfg)

	s := server.NewApiServer(cfg, h)

	log.Print("starting server...")
	log.Fatal(s.Start())
}
