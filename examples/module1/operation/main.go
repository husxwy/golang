package main

import "fmt"

func main()  {
	work1_1()
	wokr1_2()
}

//给定一个字符串数组
//["I","am","stupid","and","weak"]
//用 for 循环遍历该数组并修改为
//["I","am","smart","and","strong"]
func work1_1()  {
	mySlice := []string{"I","am","stupid","and","weak"}
	fmt.Printf("mySlice %+v\n", mySlice)
	for index, _ := range mySlice {
		// 使用下标 使用真正的指针地址
		//mySlice[index] *= 2
		if index == 2{
			mySlice[index] = "smart"
		}
		if index == 4{
			mySlice[index] = "strong"
		}
	}
	fmt.Printf("mySlice %+v\n", mySlice)

}

//- 基于 Channel 编写一个简单的单线程生产者消费者模型
//- 队列：
//队列长度10，队列元素类型为 int
//- 生产者：
//每1秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
//- 消费者：
//每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
func wokr1_2()  {

	
}