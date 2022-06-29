package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("server start🚀")
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Fatal(err)
		}
		
		go handleRequest(conn)

	}
}

func handleRequest(conn net.Conn) {
	request, err := http.ReadRequest(bufio.NewReader(conn))

	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("request.txt")
	fmt.Println("read Request")
	bytes, err := json.Marshal(request.Header)
	f.Write(bytes)
	response := http.Response{
		StatusCode: 200,
		ProtoMajor: 1,
		ProtoMinor: 0,
		Body : ioutil.NopCloser(strings.NewReader(("It works\n"))),
	}

	header := http.Header{}
	header.Add("Content-Type", "text/html")
	header.Add("Host", "webserver-of-the-ryo/0.1")
	header.Add("Date", time.Now().Format(time.UnixDate))
	header.Add("Connection", "Close")
	response.Header = header
	response.Write(conn)
	conn.Close()
}