package main

import (
	"fmt"
	"sync"
)

var (
	mux   sync.Mutex
	chain string
)

// deadlock
func main() {
	chain = "chain"
	A()
	fmt.Println(chain)
}

func A() {
	mux.Lock()
	// 这里要等到 A 返回前才能释放锁
	defer mux.Unlock()
	chain = chain + " --> A"
	B()
}

func B() {
	chain = chain + " --> B"
	C()
}

func C() {
	// 这里要等到 A 释放锁后才能获取到锁，C 无法获取锁也就无法返回，
	// 导致 A 无法返回，也就无法释放锁，
	// C 等待 A 释放锁才能返回，A 等待 C 返回才能释放锁，死锁形成
	mux.Lock()
	defer mux.Unlock()
	chain = chain + " --> C"
}
