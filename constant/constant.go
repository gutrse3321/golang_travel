package main

import "fmt"
import "unsafe"

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

const (
	e = iota // 0
	f        // 1
	g        // 2
	h = "ha" // 独立值
	i        // "ha"
	j = 100  // 独立值
	l        // 100
	m = iota // 7 恢复计数
	n        // 8
)

// 二进制想左移
const (
	pul = 1 << iota
	mui = 3 << iota // 左移1位   110        6
	mor             // 左移2位   1100      12
	moe             // 左移3位   11000     24
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Println("面积为: ", area)
	println()
	println(a, b, c)
	println(e, f, g, h, i, j, l, m, n)
	println(pul, mui, mor, moe)
}
