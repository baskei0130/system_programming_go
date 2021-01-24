package main

import (
	"archive/zip"
	"bufio"
	"crypto/rand"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

func readhttp() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	// Print header
	fmt.Println(res.Header)
	// Print body
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}

func filecopy() {
	file, err := os.Open("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file2, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer file2.Close()
	io.Copy(file2, file)
}

func createrandom() {
	file, err := os.Create("rand.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.CopyN(file, rand.Reader, 1024)
}

func writeZip() {
	// zipの内容を書き込むファイル
	file, err := os.Create("sample.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// zip file
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	// ファイルの数だけ書き込み
	a, err := zipWriter.Create("a.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(a, strings.NewReader("1つめのテキストファイル"))

	b, err := zipWriter.Create("b.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(b, strings.NewReader("2つめのテキストファイル"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment;filename=ascii_sample.zip")

	// zipファイル
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	// ファイルの数だけ書き込み
	a, err := zipWriter.Create("a.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(a, strings.NewReader("1つめのテキストファイル"))

	b, err := zipWriter.Create("b.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(b, strings.NewReader("2つめのテキストファイル"))
}

func main() {
	filecopy()
	createrandom()
	writeZip()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

//readhttp()
/*
	for {
		// 1024バイトのバッファをmakeで作る
		buffer := make([]byte, 5)
		// sizeは実際に読み込んだバイト数
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		fmt.Printf("size=%d input='%s'\n", size, string(buffer))
	}
*/
