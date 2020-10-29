/*
	使用协程和通道, 模拟线程池
*/
package main

import (
	"fmt"
	"time"
)

// worker 逻辑, 通过协程并发执行 worker
// worker 从 jobs通道接收任务, 将处理结果发送到 results通道
func PoolWorker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker", id, "processing job", job)
		// 间隔 1s 模拟worker的逻辑耗时
		time.Sleep(time.Second)
		results <- job * 2
	}
}

func main() {

	// 创建通道
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 启动 3 个worker, 现在处于阻塞状态, 因为还没有向通道发送数据
	for w := 1; w <= 3; w++ {
		go PoolWorker(w, jobs, results)
	}

	// 向 jobs 通道发送 9 条数据, 然后关闭通道
	// jobs 通道处理完成就会关闭
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// 收集所有任务的返回
	for r := 1; r <= 9; r++ {
		<-results
	}

}

/*
	执行结果
		worker 3 processing job 1
		worker 1 processing job 2
		worker 2 processing job 3
		worker 2 processing job 4
		worker 3 processing job 6
		worker 1 processing job 5
		worker 3 processing job 8
		worker 2 processing job 9
		worker 1 processing job 7

*/
