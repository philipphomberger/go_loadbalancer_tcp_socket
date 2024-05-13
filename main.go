package main

import (
	"encoding/json"
	"fmt"
	"loadbalancertcp/client"
	"loadbalancertcp/server"
	"os"
)

func main() {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := client.ServerPool{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	server.Server(configuration)
}
