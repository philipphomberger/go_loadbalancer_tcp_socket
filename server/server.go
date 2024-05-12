package server

import (
	"fmt"
	"io"
	"loadbalancertcp/client"
	"loadbalancertcp/models"
	"net"
)

var ServerPool = []string{"localhost:5432", "localhost:5433"}
var i int = 0

func Server() {
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
		i = models.Next(i, ServerPool)
		go handleConnection(conn, i)
	}
}

func handleConnection(conn net.Conn, counter int) {
	// Close the connection when we're done
	defer conn.Close()
	c := client.Client(ServerPool[counter])
	defer c.Close()

	go io.Copy(c, conn)
	io.Copy(conn, c)

	/*	// Read incoming data
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Print the incoming data
		fmt.Printf("Received: %s", buf)
		i = models.Next(i, ServerPool)
		client.Client(buf, ServerPool[i])
		if i == 0 {
			i = 1
		}*/

}
