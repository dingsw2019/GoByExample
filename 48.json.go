/**
应用 json 编码、解码
*/
package main

import (
	"encoding/json"
	"fmt"
)

// 用来演示自定义类型的编码与解码
type Response1 struct {
	Page   int
	Fruits []string
	priKey bool
}
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

// 一些编码例子
func EncodeExample() {
	// 一些原子值的例子

	// _ 为error
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB)) // true

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB)) // 1

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB)) // 2.34

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB)) // "gopher"

	// 切片转json
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB)) // ["apple","peach","pear"]

	// map 转 json
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB)) // {"apple":5,"lettuce":7}

	// 结构体转json, 编码进输出可导出的字段(首字母大写的)
	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
		priKey: false,
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B)) // {"Page":1,"Fruits":["apple","peach","pear"]}

	// 结构体转json, 字段名按配置转换成小写
	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B)) // {"page":1,"fruits":["apple","peach","pear"]}

}

// 一些解码例子
func DecodeExample() {

}

func main() {

	//EncodeExample()

	DecodeExample()
}
