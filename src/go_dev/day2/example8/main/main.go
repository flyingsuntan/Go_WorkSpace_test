package main

import "fmt"

func modify(a int) {
	a = 10
	return
}
func modefy1(a *int) {
	*a = 10

}
func main() {
	a := 5
	b := make(chan int, 1)
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	modify(a)
	fmt.Println("a=", a)
	modefy1(&a)
	fmt.Println("a=", a)
}
