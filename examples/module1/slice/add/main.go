package main

import (
	"fmt"
)

func main() {
	mySlice1 := []int{}
	mySlice1 = append(mySlice1,1)
	mySlice1 = append(mySlice1,2)
	mySlice1 = append(mySlice1,3)
	fmt.Printf("mySlice1 %+v\n", mySlice1)

}

func deleteItem(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
