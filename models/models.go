package models

import "fmt"

func Next(i int, ServerPool []string) int {
	fmt.Println(i)
	fmt.Println(ServerPool[i])
	if i >= len(ServerPool)-1 {
		return 0
	}
	return i + 1
}
