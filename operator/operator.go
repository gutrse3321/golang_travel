package main

import "fmt"

func main() {
	// +   -     *    /     %     ++     --
	// ==      !=       >      <       >=       <=
	// &&      ||           !

	// 位运算符
	// &        |       ^          <<         >>

	var a uint = 60 // 60  =  0011  1100
	var b uint = 13 // 13  =  0000  1101
	var c uint = 0

	c = a & b // 12 = 0000 1100
	fmt.Println("a & b      c的值:", c)

	c = a | b // 61 = 0011 1101
	fmt.Println("a | b        c的值:", c)

	c = a ^ b // 49 = 0011  0001
	fmt.Println("a ^ b        c的值:", c)

	c = a << 2 // 240 = 1111  0000
	fmt.Println("a << 2     c的值:", c)

	c = a >> 2 // 15 = 0000 1111
	fmt.Println("a >> 2     c的值:", c)
}
