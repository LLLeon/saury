package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mux   sync.RWMutex
	count int
)

// deadlock
func main() {
	go A()

	// 这里睡眠 2 秒后尝试获取写锁
	time.Sleep(2 * time.Second)

	// 写锁的优先级高于读锁，因为 main 尝试获取写锁的时间在程序开始运行后的 2 秒 和 5 秒之间，
	// 而 C 尝试获取读锁是在程序开始运行 5 秒之后，main 早于 C 尝试获取写锁，所以 C 在 5 秒
	// 后无法获取到读锁，C 无法返回， 那么 A 也就无法释放读锁，导致 main 无法获取写锁。
	//
	// main 等待 A 释放读锁，A 等待 C 获取/释放读锁后返回，C 由于 main 获取写锁的优先级较高而
	// 无法获取/释放读锁所以无法返回，死锁形成
	mux.Lock()
	defer mux.Unlock()
	count++
	fmt.Println(count)
}

func A() {
	// 获取读锁，只有 C 返回后才能释放读锁
	mux.RLock()
	defer mux.RUnlock()
	B()
}

func B() {
	// 这里睡眠 5 秒
	time.Sleep(5 * time.Second)
	C()
}

func C() {
	// 读锁可以并发获取，不过由于 B 中睡眠了 5 秒，所以至少要 5 秒后才会尝试获取/释放读锁
	mux.RLock()
	defer mux.RUnlock()
}
