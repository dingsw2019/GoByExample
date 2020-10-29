/*
	字符串相关函数
*/
package main

import (
	"fmt"
	s "strings"
)

// 给 fmt.Println 一个别名, 方便后面使用
var p = fmt.Println

func main() {

	// substr 是否包含在 s 中
	// Contains:        true
	p("Contains:	", s.Contains("test", "es"))

	// substr 在 s 中出现的数量
	// Count:           2
	p("Count:		", s.Count("test", "t"))

	// prefix 是否为 s 的开头字母, 必须是从头开始的
	// HasPrefix:       true
	p("HasPrefix:	", s.HasPrefix("test", "te"))

	// suffix 是否为 s 的结尾字母
	// HasSuffix:       true
	p("HasSuffix:	", s.HasSuffix("test", "st"))

	// substr 在 s 中的索引位, 索引从 0 开始
	// Index:           1
	p("Index:		", s.Index("test", "e"))

	// 将数组变成字符串, 并以 sep 作为分隔符
	// Join:            a,b
	p("Join:		", s.Join([]string{"a", "b"}, ","))

	// 计算字符串长度
	// len:             5
	p("len:		", len("hello"))

	p("Char", "hello"[1])
}
