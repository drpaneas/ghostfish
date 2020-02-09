package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	const grs = 65535 // Number of Go Routines
	var wg sync.WaitGroup
	wg.Add(grs)
	for i := 1; i <= grs; i++ {
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			// If the connection is successful close the connection and print the numberß
			if err == nil {
				conn.Close()
				fmt.Printf("TCP/%d open\n", j)
			}
		}(i)
	}
	wg.Wait() // don't return main() until all go routines are finished
}
