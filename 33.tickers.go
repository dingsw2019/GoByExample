/*
	打点器, 适用于 再某个固定时间间隔重复执行 的场景
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	// 创建打点器, 每500ms 触发一次
	ticker := time.NewTicker(time.Millisecond * 500)
	// 打点器触发的执行函数
	go func() {
		for t := range ticker.C {
			fmt.Println("tick at", t)
		}
	}()

	// 执行 1600ms 后, 停止打点器
	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("ticker stop")
}

/**
执行结果
	tick at 2020-10-28 17:48:11.681087 +0800 CST m=+0.503039833
	tick at 2020-10-28 17:48:12.182444 +0800 CST m=+1.004393974
	tick at 2020-10-28 17:48:12.679576 +0800 CST m=+1.501522838
	ticker stop

*/
