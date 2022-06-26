package main

import (
	"fmt"
	"net"
	"net/http"
	"log"
	"bufio"
	"os"
	"encoding/json"
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
		request, err := http.ReadRequest(bufio.NewReader(conn))

		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Create("request.txt")
		bytes, err := json.Marshal(request.Header)
		result, err := f.Write(bytes)

		fmt.Println(result)

	}
}