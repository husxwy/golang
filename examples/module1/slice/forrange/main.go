package main

import (
	"fmt"
)

func main() {
	// c 语言 语法
	// go 不支持指针运算
	var x int = 10
	y := &x
	fmt.Printf("指针地址: %x\n", y  )
	fmt.Printf("指针的值: %d\n", *y  )

	// 切片是连续内存并且可以动态扩展，由此引发的问题？
	// 错误用例
	a := []int{}
	fmt.Printf("a %x\n", &a)
	b := []int{1,2,3}
	fmt.Printf("b %x\n", &b)
	c := a
	fmt.Println("c ", &c)
	a = append(b, 1)
	fmt.Println("a ", &a)
	fmt.Printf("a %+v\n", a)
	fmt.Printf("b %+v\n", b)
	fmt.Printf("c %+v\n", c)

	a1 := []int{}
	b1 := []int{1,2,3}
	c1 := a1
	copy(b1,a1)

	a1 = append(a1, 2)
	copy(a1,c1)
	fmt.Printf("a1 %+v\n", a1)
	fmt.Printf("b1 %+v\n", b1)
	fmt.Printf("c1 %+v\n", c1)

	//  修改切片的值
	// 错误用例
	// go 语言全部都是值传递
	mySlice := []int{10, 20, 30, 40, 50}
	for _, value := range mySlice {
		value *= 2
	}
	fmt.Printf("mySlice %+v\n", mySlice)
	// 正确用例
	for index, _ := range mySlice {
		// 使用下标 使用真正的指针地址
		mySlice[index] *= 2
	}
	fmt.Printf("mySlice %+v\n", mySlice)
}
