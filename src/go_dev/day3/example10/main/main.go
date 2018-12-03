package main

import "fmt"

func main() {
	//LABEL1:
	//	for i :=0;i<=5;i++{
	//		for j :=0;j<=5;j++{
	//			if j == 4{
	//				continue LABEL1
	//			}
	//			fmt.Printf("i is %d and j is:%d",i,j)
	//		}
	//	}
	xh()
}

func xh() {
	str := "hello,world,中国"
	for i, v := range str {
		fmt.Printf("index[%d] val[%c] len[%d]\n", i, v, len([]byte(string(v))))
	}
}
