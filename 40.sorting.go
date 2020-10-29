/*
	内置排序函数
*/
package main

import (
	"fmt"
	"sort"
)

func main() {

	strs := []string{"c", "a", "b"}
	// 原地排序, 所以没有返回
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// 对 int 数组排序
	ints := []int{7, 6, 1}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	// 检查数组是否正序排列
	sorted := sort.IntsAreSorted(ints)
	fmt.Println("Ints Sorted:", sorted)
}

/*
	执行结果
		Strings: [a b c]
		Ints: [1 6 7]
		Ints Sorted: true
*/
