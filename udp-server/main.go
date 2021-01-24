package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Server is running at localhost:8888")
	conn, err := net.ListenPacket("udp", "localhost:8888")
	if err != nil {
		log.Print(err)
	}
	defer conn.Close()
	buffer := make([]byte, 500)
	for {
		length, remoteAddress, err := conn.ReadFrom(buffer)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("Received from %v: %v\n", remoteAddress, string(buffer[:length]))
		_, err = conn.WriteTo([]byte("Hello from Server"), remoteAddress)
		if err != nil {
			log.Print(err)
		}
	}
}
