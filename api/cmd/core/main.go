package main

import (
	"log"

	"github.com/USACE/instrumentation-api/api/internal/config"
	"github.com/USACE/instrumentation-api/api/internal/handler"
	"github.com/USACE/instrumentation-api/api/internal/server"
)

// @title MIDAS Web API
// @version 2.0
// @description Monitoring Instrumentation Data Acquisition Systems (MIDAS) Web API

// @license.name MIT
// @license.url https://github.com/USACE/instrumentation-api/blob/555ea51191ff1245fe5910a295862be7514aaec6/LICENSE.md

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and access token.

// @securityDefinitions.apikey CacOnly
// @in header
// @name Authorization
// @description CAC-Only routes
// @scope.admin Grants read and write access to administrative information
// @scope.project_admin Grants project members read and write access to projects
// @scope.project_member Read and write permissions per-project granted by project admin
func main() {
	cfg := config.NewApiConfig()

	h := handler.NewApi(cfg)

	s := server.NewApiServer(cfg, h)

	log.Print("starting server...")
	log.Fatal(s.Start())
}
