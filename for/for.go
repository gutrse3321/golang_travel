package main

import "fmt"

func main() {
	b := 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	// for 循环
	for a := 0; a < 10; a++ {
		fmt.Println("a的值为:", a)
	}

	// while一样
	for a < b {
		a++
		fmt.Println("aa的值为: ", a)
	}

	// foreach
	for i, x := range numbers {
		fmt.Println("第", i, "位x的值 = ", x)
	}

	var result int

	for n1 := 1; n1 <= 9; n1++ {
		for n2 := 1; n2 <= 9; n2++ {
			result = n1 * n2
			fmt.Printf("%d * %d = %d", n1, n2, result)
		}
	}
}
