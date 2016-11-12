package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	// Resolve TCP address and port
	udpAddr, e := net.ResolveUDPAddr("udp", "127.0.0.1:3000")
	cError(e)

	// Create TCP socket
	c, e := net.DialUDP("udp", nil, udpAddr)
	cError(e)

	// Get Current time and conver to binary
	buf, _ := time.Now().MarshalBinary()

	// Write time stamp to output stream
	c.Write(buf)
	cError(e)
}

// Helper function to test for connection errors
func cError(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", e.Error())
		os.Exit(1)
	}
}
