package main

import (
	"loadbalancertcp/client"
	"loadbalancertcp/server"
)

func main() {
	var serverPool client.ServerPool
	serverPool.Servers = append(serverPool.Servers, client.Server{Name: "localhost", Port: "5432", Available: true})
	serverPool.Servers = append(serverPool.Servers, client.Server{Name: "localhost", Port: "5433", Available: true})
	server.Server(serverPool)
}
