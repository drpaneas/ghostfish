package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup // synchronised counter for Go routines
	for i := 1; i <= 65535; i++ {
		wg.Add(1) // increment the counter each time you create a go routine to scan a port
		go func(j int) {
			wg.Done() // decrement the counter whenever a go routine has been performed
			address := fmt.Sprintf("127.0.0.1:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				// port is closed or filtered
				return
			}
			// Connection is successful
			conn.Close() // FINishing the connection is just polite
			fmt.Printf("TCP/%d open\n", j)
		}(i)
	}
	wg.Wait() // block going further until the counter has returned zero
}
