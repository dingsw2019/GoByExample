/*
	阻塞等待多个协程执行结束, 使用 sync.WaitGroup
	WaitGroup的相当于一个计数器, 执行+1 -1操作, 直到计数器为0, 程序才会退出
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

// 每个协程都会运行此函数
// 注意，WaitGroup 必须通过指针传递给函数！！！
func GroupWorker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", id)
	// 睡1s, 模拟逻辑耗时
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)

	// 通知 WaitGroup, 当前协程工作已完成, 相当于对计数器做-1
	wg.Done()
}

func main() {

	// 申请 WaitGroup
	var wg sync.WaitGroup

	// 开启多个协程, 并递增 WaitGroup 计数器
	// 相当于监听所有协程的执行状态
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go GroupWorker(i, &wg)
	}

	fmt.Println("work sent all")
	// 阻塞主程序, 直到所有协程执行完成后程序结束
	wg.Wait()
	fmt.Println("work all done")
}

/**
运行结果
	Worker 1 starting
	Worker 2 starting
	work sent all
	Worker 5 starting
	Worker 3 starting
	Worker 4 starting
	Worker 5 done
	Worker 2 done
	Worker 3 done
	Worker 1 done
	Worker 4 done
	work all done


*/
