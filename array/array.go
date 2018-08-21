package main

import "fmt"

func main() {
	var n [10]int
	var i, j int

	// 给数组初始化元素
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Println("index:", j, "item: ", n[j])
	}
}
