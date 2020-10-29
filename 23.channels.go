/**
通道, 是协程之间传递数据的通道
A协程发送数据到通道, B协程中接收数据
*/
package main

import "fmt"

func main() {

	// 创建一个传字符串的通道, 语法: make(chan val-type)
	messages := make(chan string)

	// 向通道发送数据, 语法: channel <- 发送内容
	// 以下语句, 通过协程将 ping 发到 messages通道
	go func() {
		messages <- "ping"
	}()

	// 从通道中接收数据, 语法: <- channel
	// 以下语句, 从message通道接收一个值
	msg := <-messages

	fmt.Println(msg)
}

/*
	默认发送和接收操作是阻塞的，直到发送方和接收方都准备完毕。
	这个特性允许我们，不使用任何其它的同步操作，来在程序结尾等待 消息 "ping"。

	运行结果
	ping
*/
