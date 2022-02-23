package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fullString := "hello world"
	fmt.Println(fullString)
	// i index 数组下标 c char 字符
	for i, c := range fullString {
		fmt.Println(i, string(c))
	}

}
