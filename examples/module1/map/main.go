package main

import (
	"fmt"
)

func main() {
	myMap := make(map[string]string, 10)
	myMap["a"] = "b"
	myFuncMap := map[string]func() int{
		"funcA": func() int { return 1 },
	}
	fmt.Println(myFuncMap)
	f := myFuncMap["funcA"]
	fmt.Println(f())
	// 正例 获取value 判断是否存在，存在打印，避免空指针
	value, exists := myMap["a"]
	if exists {
		println(value)
	}
	for k, v := range myMap {
		println(k, v)
	}
}
