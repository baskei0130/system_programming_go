package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn = nil
	for {
		var err error
		// まだコネクションを張っていない / エラーでリトライ時はDialから行う
		if conn == nil {
			conn, err = net.Dial("tcp", "localhost:8888")
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("Access: %d\n", current)
		}
		// POST で文字列を送るリクエストを作成
		request, err := http.NewRequest(
			"POST",
			"http://localhost:8888",
			strings.NewReader(sendMessages[current]))
		if err != nil {
			log.Print(err)
		}
		request.Header.Set("Accept-Encoding", "gzip")
		request.Write(conn)
		//conn.Write([]byte("GET / HTTP/1.0\r\nHost: localhost:8888\r\n\r\n"))
		response, err := http.ReadResponse(
			bufio.NewReader(conn), request)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}
		// 結果を表示
		dump, err := httputil.DumpResponse(response, false)
		if err != nil {
			log.Print(err)
		}
		fmt.Println(string(dump))
		defer response.Body.Close()

		if response.Header.Get("Content-Encoding") == "gzip" {
			reader, err := gzip.NewReader(response.Body)
			if err != nil {
				log.Print(err)
			}
			io.Copy(os.Stdout, reader)
			defer reader.Close()
		} else {
			io.Copy(os.Stdout, response.Body)
		}
		// 送信完了していれば終了
		current++
		if current == len(sendMessages) {
			break
		}
	}
}
