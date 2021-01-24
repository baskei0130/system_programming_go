package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Listen tick server at 224.0.0.1:9999")
	address, err := net.ResolveUDPAddr("udp", "224.0.0.1:9999")
	if err != nil {
		log.Print(err)
	}
	listener, err := net.ListenMulticastUDP("udp", nil, address)
	if err != nil {
		log.Print(err)
	}
	defer listener.Close()
	buffer := make([]byte, 1500)
	for {
		length, remoteAddress, err := listener.ReadFromUDP(buffer)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("Server %v\n", remoteAddress)
		fmt.Printf("Now    %s\n", string(buffer[:length]))
	}
}
