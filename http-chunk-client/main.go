package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func main() {
	current := 0
	var conn net.Conn = nil
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
		"GET",
		"http://localhost:8888",
		nil)
	if err != nil {
		log.Print(err)
	}
	err = request.Write(conn)
	if err != nil {
		log.Print(err)
	}
	reader := bufio.NewReader(conn)
	//conn.Write([]byte("GET / HTTP/1.0\r\nHost: localhost:8888\r\n\r\n"))
	response, err := http.ReadResponse(reader, request)
	if err != nil {
		log.Print(err)
	}
	// 結果を表示
	dump, err := httputil.DumpResponse(response, false)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(string(dump))
	if len(response.TransferEncoding) < 1 ||
		response.TransferEncoding[0] != "chunked" {
		panic("wrong transfer encoding")
	}

	for {
		// get size
		sizeStr, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		// １６進数のサイズをパース，サイズが0ならクローズ
		size, err := strconv.ParseInt(string(sizeStr[:len(sizeStr)-2]), 16, 64)
		if size == 0 {
			break
		} else if err != nil {
			log.Print(err)
		}
		// サイズ分バッファを確保して読み込み
		line := make([]byte, int(size))
		reader.Read(line)
		reader.Discard(2)
		fmt.Printf("%d bytes: %s\n", size, string(line))
	}
}
