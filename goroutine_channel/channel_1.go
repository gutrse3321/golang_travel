package main

import (
	"fmt"
	"time"
)

var c = make(chan string, 3)

func main() {
	syncC1 := make(chan struct{}, 1)
	syncC2 := make(chan struct{}, 2)
	// 接收操作
	go func() {
		// 等待接收信号，在发送方发送完这个信号之前，当前goroutine会一直等待
		// 接收到信号，等待1s，等 c 的前三个值接收完成再接收
		<-syncC1
		fmt.Println("[接收]等1秒")
		time.Sleep(time.Second)
		for {
			if elem, ok := <-c; ok {
				fmt.Println("接收: ", elem)
			} else {
				break
			}
		}
		fmt.Println("接收停止")
		syncC2 <- struct{}{}
	}()
	// 发送操作
	go func() {
		for _, elem := range []string{"a", "b", "c", "d"} {
			c <- elem
			fmt.Println("发送: ", elem)
			if elem == "c" {
				// 当遍历到第三个数的时候发送信号
				syncC1 <- struct{}{}
				fmt.Println("发送一个同步信号")
			}
		}
		fmt.Println("等待2秒")
		time.Sleep(time.Second * 2)
		close(c)
		syncC2 <- struct{}{}
	}()
	<-syncC2
	<-syncC2
}
