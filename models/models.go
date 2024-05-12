package models

import (
	"errors"
	"loadbalancertcp/client"
)

func Next(i int, ServerPool client.ServerPool) int {
	if ok, err := CheckAny(ServerPool); ok && err == nil {
		if i >= len(ServerPool.Servers)-1 {
			return 0
		}
		return i + 1
	}
	return 0
}

func CheckAny(ServerPool client.ServerPool) (bool, error) {
	counter := 0
	for _, server := range ServerPool.Servers {
		if server.Available {
			counter++
		}
	}
	if counter == 0 {
		return false, errors.New("pool: No Servers available")
	}
	return true, nil
}
