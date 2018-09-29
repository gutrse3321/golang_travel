package main

import "fmt"

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(a[:len(a) / 2], c)
	go sum(a[len(a) / 2:], c)

	// 从c获取数据，并赋值
	x, y := <-c, <-c

	// shell: 17 -5 12
	// me:洋葱式赋值
	fmt.Println(x, y, x + y)
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}

	// 发送total给c
	c <- total
}