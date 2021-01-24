package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var source = `1行目
2行目
3行目`
var source2 = "123 1.234 1.0e4 test"
var csvSource = `13101,"100  ","1000003","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ﾋﾄﾂﾊﾞｼ(1ﾁｮｳﾒ)","東京都","千代田区","一ツ橋（１丁目）",1,0,1,0,0,0
13101,"101  ","1010003","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ﾋﾄﾂﾊﾞｼ(2ﾁｮｳﾒ)","東京都","千代田区","一ツ橋（２丁目）",1,0,1,0,0,0
13101,"100  ","1000012","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ﾋﾋﾞﾔｺｳｴﾝ","東京都","千代田区","日比谷公園",0,0,0,0,0,0
13101,"102  ","1020093","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ﾋﾗｶﾜﾁｮｳ","東京都","千代田区","平河町",0,0,1,0,0,0
13101,"102  ","1020071","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ﾌｼﾞﾐ","東京都","千代田区","富士見",0,0,1,0,0,0
`

func pracReader() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%#v\n", line)
		break
	}
}

func pracScanner() {
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}

func pracFscan() {
	reader := strings.NewReader(source2)
	var i int
	var f, g float64
	var s string
	fmt.Fscan(reader, &i, &f, &g, &s)
	fmt.Printf("i=%#v f=%#v g=%#v s=%#v\n", i, f, g, s)
}

func pracCSVRead() {
	reader := strings.NewReader(csvSource)
	csvReader := csv.NewReader(reader)
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(line[2], line[6:9])
	}
}

func pracMultiReader() {
	header := bytes.NewBufferString("----HEADER----\n")
	body := bytes.NewBufferString("Example of MultiReader\n")
	footer := bytes.NewBufferString("----FOOTER---\n")

	reader := io.MultiReader(header, body, footer)
	io.Copy(os.Stdout, reader)
	reader = bufio.NewReader(strings.NewReader(source))
	copyNmade(os.Stdout, reader, 10)
}

func copyNmade(dst io.Writer, src io.Reader, length int) error {
	length64 := int64(length)
	written, err := io.Copy(dst, io.LimitReader(src, length64))
	if written == length64 {
		return nil
	}
	if written < length64 && err == nil {
		return io.EOF
	}
	return err
}

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func pazzle() {
	var stream io.Reader
	a := io.NewSectionReader(programming, 5, 1)
	s := io.LimitReader(system, 1)
	c := io.LimitReader(computer, 1)
	i := io.NewSectionReader(programming, 8, 1)
	pr, pw := io.Pipe()
	writer := io.MultiWriter(pw, pw)
	go io.CopyN(writer, i, 1)
	defer pw.Close()
	stream = io.MultiReader(a, s, c, io.LimitReader(pr, 2))
	io.Copy(os.Stdout, stream)
}

func main() {
	pracReader()
	pracScanner()
	pracFscan()
	pracCSVRead()
	pracMultiReader()
	pazzle()
}
