/*
	Unix时间的 秒数、毫秒数、微妙数
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	secs := now.Unix()      // 秒
	nanos := now.UnixNano() // 纳秒

	fmt.Println(now)   // 2020-10-30 14:37:36.297633 +0800 CST m=+0.000071389
	fmt.Println(secs)  // 1604039856
	fmt.Println(nanos) // 1604039906422773000

	// 毫秒要自己算, 这...
	millis := nanos / 1000000
	fmt.Println(millis)

	// 将秒、纳秒转成时间
	fmt.Println(time.Unix(secs, 0))  // 2020-10-30 14:41:13 +0800 CST
	fmt.Println(time.Unix(0, nanos)) // 2020-10-30 14:41:13.903703 +0800 CST
}
