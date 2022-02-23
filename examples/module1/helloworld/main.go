package main

import (
	"flag"
	"fmt"
	"os"
)
// main 程序入口函数 没有入参
func main() {
	// --name 默认参数 world
	// go
	// ./main --help
	name := flag.String("name", "world", "specify the name you want to say hi")
	flag.Parse()
	fmt.Println("os args is:", os.Args)
	fmt.Println("input parameter is:", *name)
	fullString := fmt.Sprintf("Hello %s from Go\n", *name)
	fmt.Println(fullString)
}
