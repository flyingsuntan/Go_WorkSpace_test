package main

import "fmt"

var sum int


func test_goroute(a int, b int)  {

	sum = a + b
	fmt.Println(sum)
}
func test_print(a int)  {
	fmt.Println(a)
}