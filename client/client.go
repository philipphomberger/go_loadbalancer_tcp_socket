package client

import (
	"fmt"
	"io"
	"net"
	"time"
)

type Server struct {
	Name      string
	Port      string
	Available bool
}

type ServerPool struct {
	Servers []Server
}

func GetReadyTargets(servers ServerPool) ServerPool {
	for i := 0; i < len(servers.Servers); i++ {
		if HealthCheck(servers.Servers[i].Name, servers.Servers[i].Port) {
			servers.Servers[i].Available = true
		} else {
			servers.Servers[i].Available = false
		}
	}
	return servers
}

func HealthCheck(host string, port string) bool {
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		fmt.Println("Connecting error:", err)
		return false
	}

	if conn != nil {
		defer conn.Close()
		fmt.Println("Opened", net.JoinHostPort(host, port))
		return true
	}
	return false
}

func Client(server string) io.ReadWriteCloser {
	// Connect to the server
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Println(err)
	}
	return conn

}
