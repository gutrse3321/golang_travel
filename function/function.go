package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 100
	var b int = 200

	result := max(a, b)
	str1, str2 := swap("tomo", "zzs")

	fmt.Println("最大值是: ", result)
	fmt.Println("名字是: ", str1, str2)
	fmt.Printf("交换前,a的值: %d\n", a)
	fmt.Printf("交换前, b的值: %d\n", b)
	/* 调用 callByRefer() 函数
	 * &a 指向 a 指针，a 变量的地址
	 * &b 指向 b 指针，b 变量的地址
	 */
	callByRefer(&a, &b)
	fmt.Printf("交换后，a 的值 : %d\n", a)
	fmt.Printf("交换后，b 的值 : %d\n", b)

	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))

	// nextNumber 是一个函数，函数i是0
	nextNumber := getSquence()

	// 调用nextNumber 函数，i变量自增1并返回
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	// 创建新的函数
	nextNumber1 := getSquence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}

// 函数返回两个数的最大值
func max(num1, num2 int) int {
	// 声明局部变量
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

// 返回两个参数
func swap(x, y string) (string, string) {
	return x, y
}

// 函数引用传递值
func callByRefer(x *int, y *int) {
	var temp int
	// 保存x地址上的值
	temp = *x
	// 将y值赋给x
	*x = *y
	*y = temp
}

func getSquence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
