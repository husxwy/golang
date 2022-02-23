package main

import (
	"reflect"
)

type MyType struct {
	// json:"name" 标签
	Name string `json:"name"`
}

func main() {
	mt := MyType{Name: "test"}
	// 反射机制 获取类型
	myType := reflect.TypeOf(mt)
	name := myType.Field(0)
	// 获取标签
	tag := name.Tag.Get("json")
	println(tag)
	tb := TypeB{P2: "p2", TypeA: TypeA{P1: "p1"}}
	//可以直接访问 TypeA.P1
	println(tb.P1)
}

type TypeA struct {
	P1 string
}

type TypeB struct {
	P2 string
	TypeA
}
