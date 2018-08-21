// main
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 切片
func main() {
	// 切片是一个轻量级的结构体封装
	// 1.声明切片
	scores := make([]int, 0, 10)

	// 2.使用append扩展切片
	// 声明的切片长度为0，容量为10，使用append即扩展长度
	// ↓ 接长数组问题，如果达到了上限，append会分配新更大的数组
	// 并把值复制过去，所以要将append返回的值重新赋给scores
	// 关于一次扩展多少 ↓ getMore函数
	//	scores = append(scores, 5)

	// 3.切分，给第切片的第六个元素赋值
	// 调整切片的大小上限是这个切片的容量，也就是10
	// 关于长数组问题 回到 append ↑
	scores = scores[0:6]
	scores[5] = 9033

	fmt.Println(scores) // 2.[5]
	// 3.[0 0 0 0 0 9033]

	// 4.
	// 接扩展多少 ↑ 2倍算法，x2 增加
	getMore()
	//this c: 5
	//index:  5 item c:  10
	//index:  10 item c:  20
	//index:  20 item c:  40

	//5.从乱序中去除一个值，js的splice
	group := []int{1, 2, 3, 4, 5, 6}
	group = removeAtIndex(group, 2)
	fmt.Println("group: ", group)
	// 5. group:  [1 2 6 4 5]
	// 输出可看，调整位置后移除后，顺序已乱

	//	6.copy内置函数
	// copyMethod()中使用了其他内置函数
	// 可以从输出中发现，copy的始终是声明worst切片的长度和容量的值
	copyMethod()
	// 6.[0 2 2 15 26]
}

func getMore() {
	scores := make([]int, 0, 5)

	c := cap(scores)

	fmt.Println("this c:", c)

	for i := 0; i < 25; i++ {
		scores = append(scores, i)
		// 如果容量改变，go增长数组的长度
		if cap(scores) != c {
			c = cap(scores)
			fmt.Println("index: ", i, "item c: ", c)
		}
	}
}

func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1

	source[index], source[lastIndex] = source[lastIndex], source[index]

	// 除了最后一个，其他保留
	// scores := []int{1, 2, 3, 4, 5}
	// scores = scores[:len(scores) - 1]
	return source[:lastIndex]
}

func copyMethod() {
	scores := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)
	worst := make([]int, 5)
	copy(worst, scores[:100])
	fmt.Println(worst)
}
