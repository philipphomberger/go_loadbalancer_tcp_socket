package server

import (
	"fmt"
	"io"
	"loadbalancertcp/client"
	"loadbalancertcp/models"
	"net"
)

var i int = 0

func Server(serverPool client.ServerPool) {
	// Listen for incoming connections on port 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Accept incoming connections and handle them
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Handle the connection in a new goroutine
		serverPool = client.GetReadyTargets(serverPool)
		i = models.Next(i, serverPool)
		if serverPool.Servers[i].Available {
			i = models.Next(i, serverPool)
		}
		go handleConnection(conn, i, serverPool)
	}
}

func handleConnection(conn net.Conn, counter int, serverPool client.ServerPool) {
	// Close the connection when we're done
	defer conn.Close()
	c := client.Client(serverPool.Servers[counter].Name + ":" + serverPool.Servers[counter].Port)
	defer c.Close()

	go io.Copy(c, conn)
	io.Copy(conn, c)
}
