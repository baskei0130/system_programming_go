package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
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
		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			log.Print(err)
		}
		fmt.Println(string(dump))
		// 送信完了していれば終了
		current++
		if current == len(sendMessages) {
			break
		}
	}
}
