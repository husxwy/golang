package main

import (
	"fmt"
)

func main() {
	name := "testing"
	// vet 警告 %d 数字  name 字符串
	fmt.Printf("%d\n", name)
	// vet 警告 2 个参数
	fmt.Printf("%s\n", name, name)
}
