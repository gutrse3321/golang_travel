package main

import "fmt"

func main() {
	a := 20
	// 声明指针变量
	var ip *int

	ip = &a

	fmt.Println("a的变量的地址是：", &a)

	// 指针变量的存储地址
	fmt.Println("ip指针变量存储的指针地址：", ip)

	// 使用指针方位值
	fmt.Println("*ip的值：", *ip)

	var ptr *int
	if ptr == nil {
		fmt.Println("是nil 空指针")
		fmt.Println("ptr的值：", ptr)
		fmt.Printf("ptr的值：%x", ptr)
	}
}
