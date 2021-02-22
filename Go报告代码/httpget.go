package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)

func main(){
	res, err := http.Get("http://www.baidu.com")
	checkError(err)
	data, err := ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Printf("Got:%q", string(data))
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Get:%v", err)
	}
}