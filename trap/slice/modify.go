package main

import (
	"fmt"
)

func main() {
	s := []int{9, 9, 9}

	// 直接通过 %p 取 slice 地址，返回的是其底层数组第一个元素的地址
	fmt.Printf("get s addr: %p\n", s)
	fmt.Printf("get &s addr: %p\n", &s)
	fmt.Printf("get s[0] addr: %p\n", &s[0])

	// 依然是传值，即 s 的拷贝，传的是一个新的 slice 结构，
	// 底层依然指向原来的数组。
	// 当发生扩容时，底层数组地址才会变。
	modify(s)
	fmt.Printf("New value: %d\n", s)
}

func modify(s []int) {
	fmt.Printf("[in func] get &s addr: %p\n", &s)
	fmt.Printf("[in func] get s addr: %p\n", s)

	s[0] = 1
}
