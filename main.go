package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please provide host:port")
		os.Exit(1)
	}

	fmt.Println("Resolving Address")
	udpAddr, err := net.ResolveUDPAddr("udp", os.Args[1])

	if err != nil {
		fmt.Println("Failed to Resolve UDP Address")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Setting up listener")
	conn, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		fmt.Println("Failed to setup Listener")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Start Listening")
	for {
		var buf [512]byte
		_, addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			fmt.Println("Failed to read request")
			fmt.Println(err)
			return
		}

		fmt.Print("> ", string(buf[0:]))

		_, err = conn.WriteToUDP([]byte("Hello UDP Client\n"), addr)
		if err != nil {
			fmt.Println("Failed to write to UDP")
			fmt.Println(err)
		}
	}

}
