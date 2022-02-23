package main

import (
	"fmt"
)

// Make 和 New
func main() {
	// New 返回指针地址
	mySlice1 := new([]int)
	// Make 返回第一个元素，可预设内存空间，避免未来的内存拷贝
	// type size
	mySlice2 := make([]int, 0)
	mySlice3 := make([]int, 10)
	mySlice4 := make([]int, 10, 20)
	fmt.Printf("mySlice1 %+v\n", mySlice1)
	fmt.Printf("mySlice2 %+v\n", mySlice2)
	fmt.Printf("mySlice3 %+v\n", mySlice3)
	fmt.Printf("mySlice4 %+v\n", mySlice4)
}
