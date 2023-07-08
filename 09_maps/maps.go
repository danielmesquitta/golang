//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package maps

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func displayServersInfo(serversInfo map[string]int) {
	println("Number of servers:", len(serversInfo))

	serversCountByStatus := make(map[int]int)

	for _, status := range serversInfo {
		serversCountByStatus[status] += 1
	}

	println("Number of servers by status:")
	for status, count := range serversCountByStatus {
		println(status, ":", count)
	}

	println()
}

func Maps() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serversInfo := make(map[string]int)

	for _, server := range servers {
		serversInfo[server] = Online
	}

	displayServersInfo(serversInfo)

	serversInfo["darkstar"] = Retired

	serversInfo["aiur"] = Offline

	displayServersInfo(serversInfo)

	for server := range serversInfo {
		serversInfo[server] = Maintenance
	}

	displayServersInfo(serversInfo)
}
