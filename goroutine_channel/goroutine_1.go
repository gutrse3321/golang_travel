package main

import (
	"fmt"
	"runtime"
	"time"
)

// 协成
// 十几个 goroutine 可能体现在底层就是五六个线程，Go 语言内部帮你实现了这些 goroutine 之间的内存共享
// 执行 goroutine 只需极少的栈内存(大概是 4~5 KB)，当然会根据相应的数据伸缩。
// 也正因为如此，可同时运行成千上万个并发任务。goroutine 比 thread 更易用、更高效、更轻便。
// in main()
// sleep in main()
// begin longWait()
// begin shortWait()
// z
// z
// z
// z
// z
// end of shortWait()
// end of longWait()
// end of main()
func main() {
	fmt.Println("in main()")
	go longWait()
	go shortWait()
	go say("z")
	fmt.Println("sleep in main()")

	time.Sleep(10 * 1e9)
	fmt.Println("end of main()")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		// runtime.Gosched() 表示让 CPU 把时间片让给别人,下次某个时候继续恢复执行该 goroutine。
		// 默认情况下，在 Go 1.5
		// 将标识并发系统线程个数的 runtime.GOMAXPROCS 的初始值由 1 改为了运行环境的 CPU 核数。
		runtime.Gosched()
		fmt.Println(s)
	}
}

func shortWait() {
	fmt.Println("begin shortWait()")
	time.Sleep(2 * 1e9)
	fmt.Println("end of shortWait()")
}

func longWait() {
	fmt.Println("begin longWait()")
	time.Sleep(5 * 1e9)
	fmt.Println("end of longWait()")
}