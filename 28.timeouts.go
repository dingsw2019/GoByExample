/*
	对于调用外部资源, 可以使用 协程+通道+select 实现超时控制
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	// 定义通道
	c1 := make(chan string)

	// 协程请求外部资源, 并将请求内容写入通道
	// 模拟外部资源的响应时间为 2s
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	// 通过select 控制接收请求内容
	// 如果获取外部资源返回时间在 1s 内, 执行第一个case, 输出返回结果
	// 如果获取外部资源返回时间大于 1s, 执行第二个case, 输出超时错误提示
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	/*
		因为外部资源响应时间为2s, 所以select的case time先到了

		执行结果
			timeout 1
	*/

	// 如果超时为 3s, 就能从c2接收到数据并打印了
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}

	/*
		执行结果
			result 2
	*/
}
