package main

import (
	"fmt"
)

// 返回参数是 result，首先执行 r=0 (因为是 return 0)，
// 随后在 defer 里面 r++ = 1，所以返回的 result=1
// r = 0
// defer r++
// return r -> 1
func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

// 返回参数是 r，首先执行 r=t=5，随后在 defer 里面 t=10，所以返回的还是 r=5。
// t := 5
// r = t
// defer t = t + 5
// return r -> 5
func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

// 首先执行 r=1，随后执行 defer，但 defer 函数的参数 r 是复制了一份传入的，
// 所以不影响 f3 返回的 r
// r = 1
// defer r(a new r) = r + 5
// return r -> 1
func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

// r = 0
// defer r++
// return r -> 1
func f4() (r int) {
	defer func() {
		r++
	}()
	return
}

// r = 0
// defer r(a new r)++
// return r -> 0
func f5() (r int) {
	defer func(r int) {
		r++
	}(r)
	return
}

func inner(i int) int {
	return i + 1
}

// r = 9
// inner r(a new r1) + 1 = 10
// defer r(a new r2)++ = 11
// return r -> 9
func f6() (r int) {
	defer func(r int) {
		r++
	}(inner(r))

	return 9
}

func main() {
	r1 := f1()
	fmt.Println("f1 result:", r1) // 1

	r2 := f2()
	fmt.Println("f2 result:", r2) // 5

	r3 := f3()
	fmt.Println("f3 result:", r3) // 1

	r4 := f4()
	fmt.Println("f4 result:", r4) // 1

	r5 := f5()
	fmt.Println("f5 result:", r5) // 0

	r6 := f6()
	fmt.Println("f6 result:", r6) // 9
}
