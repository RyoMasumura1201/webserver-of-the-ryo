package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main(){
	url := "http://localhost:80"
	responsive, _ := http.Get(url)

    defer responsive.Body.Close()

    byteArray, _ := ioutil.ReadAll(responsive.Body)
    fmt.Println(string(byteArray))
}