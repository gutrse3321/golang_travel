package main

import (
    "fmt"
)

// slice...调用传递参数调用变参函数
func main() {
    x := min(1, 3, 2, 0)
    fmt.Println("min is: ", x)
    slice := []int{7, 9, 3, 5, 1}
    x = min(slice...)
    fmt.Println("min in slice is: ", x)
}

func min(s ...int) int {
    if len(s) == 0 {
        return 0
    }
    min := s[0]
    for _, v := range s {
        if v < min {
            min = v
        }
    }
    return min
}