/*
	定时器, 用于在指定时间执行一次的场景
*/
package main

import (
	"fmt"
	"time"
)

func NowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func main() {

	fmt.Println(NowTime(), "begin")

	// 创建一个2秒后触发的定时器, 2秒后通过通道触发, 所以接收方要阻塞等待通道
	timer1 := time.NewTimer(time.Second * 2)
	fmt.Println(NowTime(), "created timer1")

	// 阻塞等待定时器触发
	<-timer1.C
	fmt.Println(NowTime(), "timer1 expired")

	// 取消定时器示例
	// 创建一个1秒后触发的定时器, 触发的操作在匿名函数中
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println(NowTime(), "timer2 expired")
	}()

	// 取消定时器, 因为1s还没到, 所以定时器还没有触发
	// 故不会执行匿名函数
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println(NowTime(), "timer2 stop")
	}
}

/*
	执行结果
		2020-10-28 17:38:56 begin
		2020-10-28 17:38:56 created timer1
		2020-10-28 17:38:58 timer1 expired
		2020-10-28 17:38:58 timer2 stop

*/
