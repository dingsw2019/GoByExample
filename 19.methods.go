/*
	挂在结构体上的函数称为"方法", 结构体+方法的模式实现类
	巧妙的设计, 函数就是对数据的操作, 结构体就是数据, 函数对结构体的操作, 让每个函数目的明确

	// 方法的结构
	func (接收器变量 接收器类型) 方法名(参数列表) (返回参数) {
		函数体
	}

	接收器是必须的, 接收器类型，可以是指针类型和非指针类型。但不能是 interface、point。
		指针类型的接收器, 相当于 this, 可以修改结构体的字段
		非指针类型的接收器, 相当于传值, 不影响结构体的字段
	有修改源数据或者大对象的情况, 使用指针类型。大对象如果是值传, copy一个对象太慢了, 传地址更快
*/
package main

import "fmt"

// 矩形结构体
type mRect struct {
	// 长宽
	width, height int
}

// 矩形结构体的方法, 用来计算面积
// 指针类型的接收器, 可以修改结构体的字段
func (r *mRect) area() int {
	r.width *= 2
	r.height *= 2
	return r.width * r.height
}

// 矩形结构体的方法, 用来计算周长
// 非指针类型的接收器, 只传值, 不影响结构体的字段
func (r mRect) perim() int {
	r.width *= 2
	r.height *= 2
	return 2*r.width + 2*r.height
}

func main() {

	r := mRect{width: 10, height: 5}

	/*
		执行结果
			perim: 60
			{10 5}
			area: 200
			{20 10}
		可以看到, area 是指针类型的接收器, 所以修改结构体字段成功, 会影响之后用到字段的方法
		但是 perim 方法是传值的, 所以不影响字段
	*/
	fmt.Println("r.perim:", r.perim())
	fmt.Println(r)
	fmt.Println("r.area:", r.area())
	fmt.Println(r)

	/**
	执行结果
		rp.perim: 120
		{20 10} &{20 10}
		rp.area: 800
		{40 20} &{40 20}
	rp 是指针, go 通过"自动解引用", 实现了指针也可以调用方法.
	rp 对字段的修改, 也会同步到 r 变量中
	*/
	rp := &r
	fmt.Println("rp.perim:", rp.perim())
	fmt.Println(r, rp)
	fmt.Println("rp.area:", rp.area())
	fmt.Println(r, rp)
}
