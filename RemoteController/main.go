package main

import (
	"flag"
	"fmt"
	"RemoteController/client"
	"RemoteController/server"
)

func main() {
	var client_flag = flag.Bool("c", false, "Specify if client")
	var server_flag = flag.Bool("s", false, "Specify if server")
	var port = flag.Int("p", 5959, "Port to bind to")
	var ip = flag.String("i", "127.0.0.1", "IP Address to bind to")
	flag.Parse()

	if *client_flag && *server_flag {
		println("Error: Please only specify client or server, not both")
	}

	if ! (*client_flag || *server_flag) {
		println("Error: Please specify if client or server")
	}

	if *client_flag {
		err := client.RunClient(port, ip)
		if err != nil {
			fmt.Println("Unable to run client: %s", err)
		}
	} else if *server_flag{
		err := server.RunServer(port, ip)
		if err != nil {
			fmt.Println("Unable to run server: %s", err)
		} }
}
