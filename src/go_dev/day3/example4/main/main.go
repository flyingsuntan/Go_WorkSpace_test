package main

import (
	"fmt"
	"time"
)

func test() {
	time.Sleep(time.Second * 5)
}
func main() {
	a := time.Now().Unix()
	fmt.Println(time.Now())
	fmt.Println(time.Now().Format("2006/1/02 15:04:05"))
	test()
	fmt.Println(time.Now().Unix() - a)
}
