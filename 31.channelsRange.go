/*
	通道遍历
*/
package main

import "fmt"

func main() {

	// 向 queue通道中写入两个值
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// 遍历 queue 通道中的值
	for elem := range queue {
		fmt.Println(elem)
	}
}

/*
	执行结果
		one
		two
*/
