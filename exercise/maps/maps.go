//--Summary:
//  Write a program to display server status.
//
//--Requirements:

//* Store the existing slice of servers in a map
//* Default all of the servers to `Online`
//* Perform the following status changes and display server info:
//  - display server info
//  - change `darkstar` to `Retired`
//  - change `aiur` to `Offline
//  - display server info
//  - change all servers to `Maintenance`
//  - display server info

package main

import (
	"fmt"
)

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

// * Create a function to print server status, including:
//   - Number of servers
//   - Number of servers for each status (Online, Offline, Maintenance, Retired)
func printServer(servers map[string]int) {
	fmt.Println("\nThere are", len(servers), "servers")

	stats := make(map[int]int)

	for _, status := range servers {
		switch status {
		case Online:
			stats[Online] += 1
		case Offline:
			stats[Offline] += 1
		case Maintenance:
			stats[Maintenance] += 1
		case Retired:
			stats[Retired] += 1

		default:
			panic("unhandled server status")

		}

	}

	fmt.Println(stats[Online], "servers are online")
	fmt.Println(stats[Offline], "servers are offline")
	fmt.Println(stats[Maintenance], "servers are undergoing maintainance")
	fmt.Println(stats[Retired], "servers are retired")
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverList := make(map[string]int)

	for _, server := range servers {
		serverList[server] = Online
	}

	printServer(serverList)

	serverList["darkstar"] = Retired
	serverList["aiur"] = Offline

	printServer(serverList)

	for server, _ := range serverList {
		serverList[server] = Maintenance
	}
	printServer(serverList)

}
