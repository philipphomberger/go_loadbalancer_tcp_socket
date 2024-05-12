package client

import (
	"fmt"
	"io"
	"net"
)

func Client(server string) io.ReadWriteCloser {
	// Connect to the server
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Println(err)
	}
	return conn

}
