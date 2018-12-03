package main

import "fmt"

func main() {

	var a int = 10

	switch a {
	case 0:
		fmt.Println("a is equal 0")
	case 10:
		fmt.Println("a is equal 10")
		//fallthrough 继续向下走
	default:
		fmt.Println("hahaha")
	}
}
