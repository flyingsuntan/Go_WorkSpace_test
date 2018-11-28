package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}
func main() {
	//var b = []int  {1,2,3,47}
	//fmt.Print(b[3])
	/*
	var n [10]int
	var i,j int

	/*为数组n初始化元数*/
	/*
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
*/
	// var a int = 10 
	// var ip *int
	// ip = &a

	// fmt.Printf("变量的地址：%x\n",&a)
	// fmt.Printf("ip变量存储的指针地址：%x\n",ip)
	// fmt.Printf("*ip变量的值：%d\n",*ip)

	// fmt.Println(Books{"Go 语言","www.runoob.com","Go 语言教程",6495407})
	// fmt.Println(Books{title:"Go 语言",author:"www.runoob.com",subject:"Go 语言教程",book_id:6495407})
	// fmt.Println(Books{title:"Go 语言",author:"www.runoob.com"})
	// s := [] int{1,2,3}
	// s := arr[:]
	nums := []int{2,3,4}
	sum := 0
	for _,num := range nums{
		sum +=num
	}

	fmt.Println("sum:",sum)

	for i,num :=range nums {
		if num ==3 {
			fmt.Println("index:",i)
		}
	}

	kvs := map[string]string{"a":"apple","b":"banana"}
	for k,v := range kvs{
		fmt.Printf("%s -> %s\n",k, v)
	}

	for i,c := range "go" {
		fmt.Println(i,c)
	}
	

}