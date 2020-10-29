/**
接口是方法的集合, 多用于抽象方法(两种类共同存在的方法),
比如矩形和圆形都有计算面积和周长的方法, 那么面积和周长就是抽象方法, 形成一个接口
以下代码实现面积和周长的接口
*/
package main

import (
	"fmt"
	"math"
)

// 定义几何体接口
type geometry interface {
	area() float64  // 计算面积
	perim() float64 // 计算周长
}

// 矩形结构体
type rect struct {
	// 长宽
	width, height float64
}

// 圆形结构体
type circle struct {
	// 半径
	raduis float64
}

// 计算矩形面积
func (r rect) area() float64 {
	return r.width * r.height
}

// 计算矩形周长
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// 计算圆形面积
func (c circle) area() float64 {
	return math.Pi * c.raduis * c.raduis
}

// 计算圆形周长
func (c circle) perim() float64 {
	return 2 * math.Pi * c.raduis
}

// 如果变量是接口类型, go会调用变量的接口方法
// 如果 g 传入的是 rect, 就调用 rect 的area和perim方法
// 如果 g 传入的是 circle 就调用的 circle 的area和perim方法
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {

	// 生成矩形
	r := rect{width: 3, height: 4}

	// 生成圆形
	c := circle{raduis: 5}

	// 打印矩形的面积和周长
	measure(r)

	// 打印圆形的面积和周长
	measure(c)
}

// 执行结果
// {3 4}
// 12
// 14
// {5}
// 78.53981633974483
// 31.41592653589793
