/**
应用 json 编码、解码
*/
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
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

	// 一个 json 数据
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// 一个存放转换 json 数据的结构
	// string 为键, interface代表值为任意类型
	var dat map[string]interface{}

	// 错误检查
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	fmt.Println(dat) // map[num:6.13 strs:[a b]]

	num1 := dat["num"]
	fmt.Println(num1, reflect.TypeOf(num1)) // 6.13 float64

	num2 := dat["num"].(float64)
	fmt.Println(num2, reflect.TypeOf(num2)) // 6.13 float64

	strs := dat["strs"].([]interface{})
	fmt.Println(strs, reflect.TypeOf(strs))

	str1 := strs[0]
	fmt.Println(str1, reflect.TypeOf(str1)) // a string

	// json 放入结构体中
	str := `{"page":1, "fruits":["apple","peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)           // &{1 [apple peach]}
	fmt.Println(res.Fruits[0]) // apple

	// Stream Json, 可以将 JSON 编码输出到 os.Writer流中, 或者作为 HTTP 响应体
	// 也可以从 HTTP 请求体中解码 json
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d) // {"apple":5,"lettuce":7}

}

func main() {

	//EncodeExample()

	DecodeExample()
}
