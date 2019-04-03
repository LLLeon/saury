package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func main() {
	var stus []student

	stus = []student{
		{Name: "one", Age: 18},
		{Name: "two", Age: 19},
	}

	data := make(map[int]*student)

	// 用for range 来遍历数组或 map 时，遍历出来的值的地址是不变的，
	// 每次遍历仅执行 struct 值的拷贝,
	// 即下面的 i, v 地址不变，遍历到第二个结构时就覆盖了第一个的值
	for i, v := range stus {
		fmt.Printf("i ->%p\n", &i)
		fmt.Printf("v ->%p\n", &v)
		data[i] = &v // 应改为：data[i] = &stus[i]
	}

	for i, v := range data {
		fmt.Printf("key=%d, value=%v \n", i, v)
	}
}
