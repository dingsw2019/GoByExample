/*
	通道选择器, 使用 select 实现, 语法结构
	select {
		case communication clause  :
		   statement(s);
		case communication clause  :
		   statement(s);
		// 你可以定义任意数量的 case
		default : // 可选
			statement(s);
	}

	select 语法的执行方式
	 - 每个 case 都是一个通信(channel)
	 - 某个通信数据到了, 就可以运行了, select 只运行一个case。其他的case就不会执行了
	 - 如果有多个 case 可运行, select 会随机选择一个case执行
	 - 如果没有 case 可运行会分成两种情况
		1. 有 default 子句, 执行它
		2. 没有 default 子句, 阻塞等待通信数据到来
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	// 创建两个通道
	c1 := make(chan string)
	c2 := make(chan string)

	// 先后向两个通道发送数据
	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second)
		c2 <- "two"
	}()

	// 开启两个监听 c1、c2 通道
	// 因为select 没有default, 所以会阻塞等待, 直到通道发来数据
	// 用 for 是因为, 要启动两个 select 来监听
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

/*
	运行结果
		received one
		received two

*/
