/*
	通道同步，可以使用通道来同步协程的执行状态
	以下代码使用阻塞的接收方式等待协程运行结束
*/
package main

import (
	"fmt"
	"time"
)

// 通过协程运行此函数, done通道用来通知协程结束完毕
func worker(done chan bool) {
	fmt.Print("working...")
	// 等待 1s
	time.Sleep(time.Second)
	fmt.Println("done")

	// 发送一个值来通知, 函数运行完成
	done <- true
}

func main() {

	// 创建一个布尔值的通道
	done := make(chan bool, 1)

	// 协程运行函数
	go worker(done)

	// 阻塞等待, 直到done通道发来值
	<-done
}

/*
	如果把 <-done 去掉, 可能在worker未运行, 程序就结束了。因为 worker 是协程异步执行的
	运行结果
		working...done
*/
