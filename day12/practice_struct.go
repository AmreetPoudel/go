package struct1

import "fmt"

func main() {

	//task1
	//Define a struct Server with fields: Name, IP, Uptime (in hours).
	// Create 2 server instances.
	// Print their info using fmt.Printf with formatting.
	type Server struct {
		Name   string
		IP     string
		Uptime int
	}
	server1 := Server{
		Name:   "backend",
		IP:     "10.20.0.99",
		Uptime: 36000,
	}
	server2 := Server{
		Name:   "prod-db",
		IP:     "10.20.0.101",
		Uptime: 36000,
	}

	fmt.Printf("Name:%s\tIP:%s\tUptime:%d\n", server1.Name, server1.IP, server1.Uptime)
	fmt.Printf("Name:%s\tIP:%s\tUptime:%d\n", server2.Name, server2.IP, server2.Uptime)

}
