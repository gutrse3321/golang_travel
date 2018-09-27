package main

import (
    "fmt"
)

// 标签配合continue可以直接跳出到LABEL位置开始
func main() {
    LABEL:
        for i := 0; i <= 5; i++ {
            for j := 0; j <= 5; j++ {
                if j == 4 {
                    continue LABEL
                }
                fmt.Println("i is ", i, " j is ", j)
            }
        }
}