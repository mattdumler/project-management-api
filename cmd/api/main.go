package main

import (
	"github.com/mattdumler/project-management/internal/routes"
	"github.com/mattdumler/project-management/internal/server"
)

func main() {
	server := server.NewServer()

	routes.RegisterRoutes(server)

	server.Run()
}
