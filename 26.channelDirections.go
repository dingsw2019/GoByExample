/*
	通道方向, 当通道作为函数的参数时, 可以指定通道只接收或只发送

	只发送的通道
	func 函数名(通道变量名 chan<- string){}

	只接收的通道
	func 函数名(通道变量名 <-chan string){}


*/
package main

import "fmt"

// 参数 pings 是只发送数据的通道, 将数据发送到 pings 通道
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 参数 pings 是只接收数据的通道
// 参数 pongs 是只发送数据的通道
// 将 pings 中接收的数据转发到 pongs 通道
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {

	// 创建发送和接收通道, 均可缓存一条数据
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	// 从 pings 通道发送消息
	ping(pings, "passed message")

	// 从 pongs 通道接收消息
	pong(pings, pongs)

	// 打印接收到的数据
	fmt.Println(<-pongs)
}

/*
	运行结果
		passed message
*/
