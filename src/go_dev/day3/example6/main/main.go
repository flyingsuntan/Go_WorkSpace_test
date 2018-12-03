package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n string
	fmt.Scanf("%s", &n)
	a, err := strconv.Atoi(n)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a)
	}
}
