package main

import "fmt"

func main() {
	var tempBreak int = 10

	// break 跳出循环
	for tempBreak < 20 {
		fmt.Println("tempBreak的值为: ", tempBreak)
		tempBreak++
		if tempBreak > 15 {
			break
		}
	}

	// continue 跳出本次循环
	tempContinue := 10
	for tempContinue < 20 {
		if tempContinue == 15 {
			tempContinue += 1
			continue
		}
		fmt.Println("tempContinue的值为: ", tempContinue)
		tempContinue++
	}
}
