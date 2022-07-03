package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
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
	data :=make([]byte, 1024)
	count, _:= conn.Read(data)
	// fmt.Println(string(data[:count]))
	request := string(data[:count])

	requestElementList := splitRequest(request)
	fmt.Println(requestElementList[0])
	content := "It works\n"
	response := http.Response{
		StatusCode: 200,
		ProtoMajor: 1,
		ProtoMinor: 0,
		ContentLength: int64(len(content)),
		Body : ioutil.NopCloser(strings.NewReader((content))),
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

func splitRequest(request string)([]string){
	reg := "\r\n|\n"

	requestElementList := regexp.MustCompile(reg).Split(request, -1)

	return requestElementList
}