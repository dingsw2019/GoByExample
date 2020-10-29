/*
	结构体是一组field的组合, 可非常方便的定义事物
*/
package main

import "fmt"

// 定义结构体
type person struct {
	name string
	age  int
}

/*
	输出结果
		{Bob 20}
		{Alice 30}
		{Fred 0}
		&{Ann 40}
		Sean
		50
		51
		s.age=51, sp.age=51
*/
func main() {

	// 初始化结构体
	fmt.Println(person{"Bob", 20})

	// 按key初始化
	fmt.Println(person{name: "Alice", age: 30})

	// 省略的字段将被赋值为零("", 0, 0.0, false等)
	fmt.Println(person{name: "Fred"})

	// 结构体指针
	fmt.Println(&person{name: "Ann", age: 40})

	// 使用 " . " 来访问结构体的字段
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// 结构体指针也可以使用 " . " 访问字段, 可以这样访问是因为 go 指针会 "自动解引用"
	// 手动解除引用的实现 (*sp).age
	sp := &s
	fmt.Println(sp.age)

	// 地址引用, 所以修改 sp, s也跟着改了
	sp.age = 51
	fmt.Printf("s.age=%d, sp.age=%d\n", s.age, sp.age)
}
