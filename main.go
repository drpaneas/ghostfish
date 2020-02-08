package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// port is closed or filtered
			continue
		}
		// Connection is successful
		conn.Close() // FINishing the connection is just polite
		fmt.Printf("TCP/%d open\n", i)
	}
}
