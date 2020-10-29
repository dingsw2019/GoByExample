/**
错误处理
go 一般情况下, 函数返回的最后一个参数为 error 类型, 调用方判断最后一个参数来确定函数是否正确执行
*/
package main

import (
	"errors"
	"fmt"
)

// 错误在最后一个返回值
func work1(age int) (int, error) {

	maxAge := 35
	if age >= maxAge {
		return age, errors.New("超过35岁的不要")
	}

	// 返回错误值为 nil, 代表没有错误
	return age, nil
}

// 通过实现 Error 方法来自定义 error 类型
// 通过自定义错误类型来表示错误信息
type ageError struct {
	age int
	msg string
}

func (e *ageError) Error() string {
	return fmt.Sprintf("%d - %s", e.age, e.msg)
}

func work2(age int) (int, error) {
	maxAge := 35
	if age >= maxAge {
		return age, &ageError{age, "超过35岁的不要"}
	}

	return age, nil
}

func main() {
	// 循环数组[20,26]
	for _, age := range []int{20, 36} {

		// 将 work1 的两个返回值分别赋给 r,e
		// 并判断 e 不为 nil, 说明报错了
		if r, e := work1(age); e != nil {
			fmt.Println("work1 failed:", e)
		} else {
			fmt.Println("work1 worked:", r)
		}
	}
	/**
	执行结果
		work1 worked: 20
		work1 failed: 超过35岁的不要
	*/

	for _, age := range []int{20, 36} {
		if r, e := work2(age); e != nil {
			fmt.Println("work2 failed:", e)
		} else {
			fmt.Println("work2 worked:", r)
		}
	}
	/**
	执行结果
		work2 worked: 20
		work2 failed: 36 - 超过35岁的不要
	*/

	// 如果要在程序中得到自定义类型的数据
	_, e := work2(36)
	if ae, ok := e.(*ageError); ok {
		fmt.Println(e)
		fmt.Println(ok)
		fmt.Println(ae.age)
		fmt.Println(ae.msg)
		/*
			执行结果
				36 - 超过35岁的不要
				true
				36
				超过35岁的不要
		*/
	}
}
