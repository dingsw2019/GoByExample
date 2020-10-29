/*
	关闭通道, 意味着不能在向通道发送数据, 但是可以从通道接收数据
*/
package main

import "fmt"

func main() {

	// jobs 通道传递数据
	// done 通道阻塞程序, 直到done传递数据再结束
	jobs := make(chan int, 5)
	done := make(chan bool)

	// 死循环从 jobs 通道读取数据,
	// 如果通道关闭且通道内无数据了, more的值为false, 否则是true
	go func() {
		for {
			job, more := <-jobs
			if more {
				fmt.Println("received job", job)
			} else {
				fmt.Println("received all jobs")
				// 向done通道发送数据, 来结束主程序
				done <- true
				// 结束协程运行
				return
			}
		}
	}()

	// 向jobs通道发送数据, 然后关闭通道
	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("sent job", i)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// 匿名函数将jobs通道内数据全部接收处理完成后
	// 会向 done通道发送数据, 来结束程序执行
	<-done
}

/*
	运行结果
		sent job 1
		sent job 2
		sent job 3
		sent all jobs
		received job 1
		received job 2
		received job 3
		received all jobs

*/
