package main

import (
	"fmt"
	"sync"
	"time"
)

// panic: sync: WaitGroup is reused before previous Wait has returned
func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		// 这里在 Wait 返回前又重用了 wg
		wg.Add(1)
	}()

	fmt.Println("hello")

	wg.Wait()
}
