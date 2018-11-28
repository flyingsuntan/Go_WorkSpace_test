package main

import "fmt"

func add(a int, b int,c chan int)  {
	sum := a + b
	c <- sum
}
func main() {
	var pipe chan int
	pipe = make(chan int,1)
	go add(2,5,pipe)
	sum  := <- pipe
	fmt.Println(sum)
	//fmt.Println("Hello Word!")

	//var c = add(3,4)
	//go test_goroute(200,300)
	//fmt.Println(c)
	/*for i :=0; i < 100; i++ {
		go test_print(i)
	}
	time.Sleep(10*time.Second)*/
	//fmt.Println("start goroute")
	//test_pipe()
	//fmt.Println("end goroute")
	//time.Sleep(10*time.Second)
}