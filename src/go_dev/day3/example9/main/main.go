package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	for i := 1; i <= 5; i++ {
		a += strings.Repeat("A", i) + "\n"
	}
	fmt.Println(a)
}
