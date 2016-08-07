package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s ip_address\n", os.Args[0])
		return
	}

	addr := net.ParseIP(os.Args[1])

	if addr == nil {
		fmt.Printf("Invalid address: %s\n", os.Args[1])
	} else {
		fmt.Printf("IP address: %s\n", addr)
	}
}
