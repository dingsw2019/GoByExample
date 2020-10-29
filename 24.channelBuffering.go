/*
	通道缓冲区, 默认通道是无缓冲区的。
	 - 无缓冲区时,延迟发送。接收通道准备好后, 才能发送
	 - 有缓冲区时,实时发送。允许无接收方下, 缓存一定数量的值。
*/
package main

import "fmt"

func main() {

	// 创建一个字符串通道，最多缓冲 2个字符串
	messages := make(chan string, 2)

	// 发送数据到通道, 没有接收方, 数据被缓存起来
	messages <- "buffered"
	messages <- "channel"

	// 接收通道数据
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

/*
	运行结果
		buffered
		channel
*/
