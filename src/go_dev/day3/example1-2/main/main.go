package main

import (
	"fmt"
	"strings"
)

var url string = "192.168.1.191"

func main() {
	if strings.HasPrefix(url, "http://") {
		fmt.Println("yes")
	} else {
		fmt.Printf("http://%s\n", url)
	}
	if strings.HasSuffix(url, "/") {
		fmt.Println("yes")
	} else {
		fmt.Printf("%s/", url)
	}
	w := strings.Index(url, "1")
	fmt.Println(w)
	lw := strings.LastIndex(url, "1")
	fmt.Println(lw)

	th := "heheheword"
	b := strings.Replace(th, "he", "wo", -1)
	fmt.Println(b)

}
