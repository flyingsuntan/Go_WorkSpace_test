package main

import "fmt"

func calc(a int, b int) (int, int)  {
	c := a +b
	d := (a + b) / 2
	return  c,d
}
func main() {
	sum,avg := calc(3,4)
	fmt.Println(sum,avg)
}