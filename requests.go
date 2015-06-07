package main

import (
	// "fmt"
	"github.com/franela/goreq"
)

func main() {
	res, _ := goreq.Request{Uri: "http://ip.jsontest.com/"}.Do()
	// fmt.Println(res.Body.ToString(), err)
	// fmt.Println(err)
	body, _ := res.Body.ToString()
	// fmt.Println(res.Header.Get("Content-Type"))
	// fmt.Println(res.Body.ToString())
	println(body)

}

