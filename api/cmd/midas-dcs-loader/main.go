package main

import (
	"log"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/handler"
)

func main() {
	cfg := config.NewDcsLoaderConfig()

	h := handler.NewDcsLoader(cfg)

	log.Print("listening on queue...")
	log.Fatal(h.Start())
}
