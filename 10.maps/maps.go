package main

import (
	"fmt"
)

type ServerStatus map[string]int

const (
	Online = iota
	Offline
	Maintenance
	Retired
)

func printServerInfo(serverStatus ServerStatus) {
	var online, offline, maintenance, retired int

	for _, status := range serverStatus {
		switch status {
		case Online:
			online++
		case Offline:
			offline++
		case Maintenance:
			maintenance++
		case Retired:
			retired++
		}
	}

	fmt.Println("Server Status: ")
	fmt.Println("Online: ", online)
	fmt.Println("Offline: ", offline)
	fmt.Println("Maintenance: ", maintenance)
	fmt.Println("Retired: ", retired)
	fmt.Println("")
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverStatus := make(ServerStatus)

	for _, server := range servers {
		serverStatus[server] = Online
	}

	printServerInfo(serverStatus)

	serverStatus["darkstar"] = Retired
	serverStatus["aiur"] = Offline

	printServerInfo(serverStatus)

	for server := range serverStatus {
		serverStatus[server] = Maintenance
	}

	printServerInfo(serverStatus)
}
