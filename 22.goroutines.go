/**
协程 是轻量级的线程, 并发利器
*/
package main

import (
	"fmt"
)

// 由协程来执行此函数
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// 同步调用函数f
	f("direct")

	// 在一个 go协程中调用函数f
	// 协程间是并发执行的
	go f("goroutine")

	// 协程调用匿名函数
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 两个协程独立异步执行, 所以 26～31 行在go协程中执行
	// 程序直接运行到这一行
	fmt.Scanln() // 按任意键退出
	fmt.Println("done")
}

/*
	当我们运行这个程序时，将首先看到阻塞式调用的输出，
	然后是两个 Go 协程的交替输出。这种交替的情况表示 Go 运行时是以并发的方式运行协程的。

	运行结果, goroutine、going 异步交替执行, <enter>表示按任意键退出
		direct : 0
		direct : 1
		direct : 2
		goroutine : 0
		going
		goroutine : 1
		goroutine : 2
		<enter>
		done
*/
