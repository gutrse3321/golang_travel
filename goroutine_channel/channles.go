package main

import "fmt"

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	// 初始化通道
	// 元素类型是int的通道类型
	// 还有make(chan int, 10) 第二个参数，表示该通道在同一时间最多可以缓冲10个元素值
	// 如果省略，则表示该通道永远无法缓冲任何元素值，发送给它的元素值应该被立刻取走
	c := make(chan int)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	// 从c获取数据，并赋值
	x, y := <-c, <-c

	// shell: 17 -5 12
	// me:洋葱式赋值
	fmt.Println(x, y, x+y)
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}

	// 发送total给c
	c <- total
}
