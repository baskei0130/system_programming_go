package main

import (
	"os"
	"bytes"
	"fmt"
	//"net"
	"net/http"
	"io"
	"time"
	"encoding/json"
)

func handler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello": "world",
	})
	w.Write([]byte("http.ResponsWriter example\n"))
}

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File example\n"))
	file.Close()
	os.Stdout.Write([]byte("os.Stdout example\n"))
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	buffer.WriteString("bytes.Buffer example\n")
	fmt.Println(buffer.String())

	/*
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	io.Copy(os.Stdout, conn)
	*/

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

	file2, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file2, os.Stdout)
	io.WriteString(writer, "io.MultiWriter example\n")
	fmt.Fprintf(os.Stdout, "Write with os.Stdout ad %v", time.Now())
}
