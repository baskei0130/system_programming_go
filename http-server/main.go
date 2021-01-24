package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Print(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}
		go func() {
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			for {
				// Config Timeout
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))
				// Read Requst
				request, err := http.ReadRequest(bufio.NewReader(conn))
				if err != nil {
					// タイムアウトもしくはソケットクローズ時は終了
					// それ以外はエラー
					neterr, ok := err.(net.Error)
					if ok && neterr.Timeout() {
						fmt.Println("Timeout")
						break
					} else if err == io.EOF {
						break
					}
					log.Print(err)
				}
				dump, err := httputil.DumpRequest(request, true)
				if err != nil {
					log.Print(err)
				}
				fmt.Println(string(dump))
				content := "Hello World\n"
				// Write Response
				response := http.Response{
					Status:        "200 OK",
					StatusCode:    200,
					ProtoMajor:    1,
					ProtoMinor:    1,
					ContentLength: int64(len(content)),
					Body:          ioutil.NopCloser(strings.NewReader("Hello World\n")),
				}
				response.Write(conn)
			}
			conn.Close()
		}()
	}
}
