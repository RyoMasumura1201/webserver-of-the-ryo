package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

func main(){
	url := "localhost:80"

	conn, err := net.Dial("tcp", url)
	if err != nil {
		log.Fatal(err)
	}

	request, err:= http.NewRequest("GET", "http://localhost:80", nil)
	request.Write(conn)

	response, err := http.ReadResponse(bufio.NewReader(conn), request)

    byteArray, _ := ioutil.ReadAll(response.Body)
    fmt.Println(string(byteArray))
	response.Body.Close()
	conn.Close()
}