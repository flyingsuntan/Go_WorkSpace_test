package main

import "fmt"

func main() {
	var countryCaptalMap map[string]string /*创建集合*/
	countryCaptalMap = make(map[string]string)

	/* map插入key-value对，各个国家对应的首都*/
	countryCaptalMap["France"] = "Paris"
	countryCaptalMap["Italy"] = "罗马"
	countryCaptalMap["Japan"] = "东京"
	countryCaptalMap["India"] = "新德里"
	for country := range countryCaptalMap {
		fmt.Println(country,"首都是",countryCaptalMap[country])
	}

	captial,ok :=countryCaptalMap["美国"]
	if(ok){
		fmt.Println("美国的首都是",captial)
	}else{
		fmt.Println("美国的首都不存在")
	}
}