package main

import "fmt"

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp

}
func main() {
	first := 100
	second := 200
	//swap(&first, &second)
	first, second = second, first
	fmt.Println("first=", first)
	fmt.Println("second=", second)
}
