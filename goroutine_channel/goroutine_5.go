package main

import (
	"fmt"
	"time"
)

func main() {
	names := []string{"fuck", "shit", "mother", "fucker"}
	for _, name := range names {
		go func(who string) {
			fmt.Printf("Hello, %s\n", who)
		}(name)
	}
	time.Sleep(time.Microsecond)
}
