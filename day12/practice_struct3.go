package main

import "fmt"

type Server struct {
	Name   string
	IP     string
	Uptime int
	Status string
}

func create_server(name string, ip string, uptime int, status string) *Server {
	return &Server{
		Name:   name,
		IP:     ip,
		Uptime: uptime,
		Status: status,
	}

}

func main() {

	// Expand the Server struct with a Status field (string: "up", "down").
	// Add method Start() → sets Status to "up"
	// Add method Shutdown() → sets Status to "down
	// Add method Info() → returns formatted string summary like "Server: DB01, Status: up, IP: 10.0.0.1"

	// server1 := Server{
	// 	Name:   "backend",
	// 	IP:     "10.20.0.99",
	// 	Uptime: 36000,
	// 	Status: "up",
	// }
	// server2 := Server{
	// 	Name:   "prod-db",
	// 	IP:     "10.20.0.101",
	// 	Uptime: 36000,
	// 	// Status: "down",

	// }
	server1 := create_server("backend", "10.20.0.99", 36000, "up")
	server2 := create_server("backend", "10.20.0.101", 36000, "down")

	server1_status_before := server1.Info()
	server2_status_before := server2.Info()
	fmt.Printf("before\n%s\n%s", server1_status_before, server2_status_before)

	server1.Shutdown()
	server2.Start()

	server1_status_after := server1.Info()
	server2_status_after := server2.Info()
	fmt.Printf("after\n%s\n%s", server1_status_after, server2_status_after)

}

func (structptr *Server) Start() {
	structptr.Status = "up"
}

func (structptr *Server) Shutdown() {
	structptr.Status = "down"
}

func (server_instance Server) Info() string {
	info := fmt.Sprintf("Name:%s\tIP:%s\tUptime:%d\tStatus:%s\n", server_instance.Name, server_instance.IP, server_instance.Uptime, server_instance.Status)
	return info

}
