package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// Resolve TCP address and port
	udpAddr, e := net.ResolveUDPAddr("udp", ":3000")
	cError(e)

	// Listen on all interfaces for incoming TCP connections
	l, e := net.ListenUDP("udp", udpAddr)
	cError(e)
	defer l.Close()

	for {
		var b = make([]byte, 512)
		n, _, e := l.ReadFrom(b)
		if e != nil {
			continue
		}
		t := new(time.Time)
		t.UnmarshalBinary(b[:n])

		fmt.Printf("Received Packet at: %s\n", t.String())

		tt := time.Now().Sub(*t)
		fmt.Printf("Trip Time: %+v\n", tt)
	}
}

// Helper function to test for connection errors
func cError(e error) {
	if e != nil {
		fmt.Printf("Error: %s", e.Error())
		os.Exit(1)
	}
}
