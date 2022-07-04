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

const STATIC_PATH ="static"

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
	fmt.Println(string(data[:count]))
	request := string(data[:count])

	requestElementList := splitRequest(request)
	requestLine := requestElementList[0]
	fmt.Println(requestElementList[0])

	_, path, _ := splitRequestLine(requestLine)
	fmt.Println(path)

	responseContents := getResponseContents(path)

	fmt.Println(responseContents)

	response := makeResponse(responseContents)
	
	response.Write(conn)
	conn.Close()
}

func splitRequest(request string)([]string){
	reg := "\r\n|\n"

	requestElementList := regexp.MustCompile(reg).Split(request, -1)

	return requestElementList
}

func splitRequestLine(requestLine string)(method, path, version string){
	requestLineList := strings.Split(requestLine, " ")
	method = requestLineList[0]
	path = requestLineList[1]
	version = requestLineList[2]
	return
}

func getResponseContents(path string)(string){
	contents, err := ioutil.ReadFile(STATIC_PATH + path)

	if err != nil {
		panic(err)
	}

	return string(contents)
}

func makeResponse(responseContents string)(http.Response){
	response := http.Response{
		StatusCode: 200,
		ProtoMajor: 1,
		ProtoMinor: 0,
		ContentLength: int64(len(responseContents)),
		Body : ioutil.NopCloser(strings.NewReader((responseContents))),
	}

	header := http.Header{}
	header.Add("Content-Type", "text/html")
	header.Add("Host", "webserver-of-the-ryo/0.1")
	header.Add("Date", time.Now().Format(time.UnixDate))
	header.Add("Connection", "Close")
	response.Header = header

	return response
}