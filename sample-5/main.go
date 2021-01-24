package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		log.Print(err)
	}
	defer file.Close()
	file.Write([]byte("system call example\n"))
}
