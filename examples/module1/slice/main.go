package main

import (
	"fmt"
)

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}
	mySlice := myArray[1:3]
	fmt.Printf("mySlice %+v\n", mySlice)
	fullSlice := myArray[:]
	// 没有原生的删除方法
	remove3rdItem := deleteItem(fullSlice, 2)
	fmt.Printf("remove3rdItem %+v\n", remove3rdItem)
}

// 定义删除的方法
func deleteItem(slice []int, index int) []int {
	// 元素之前的切片和 元素之后的切片
	return append(slice[:index], slice[index+1:]...)
}
