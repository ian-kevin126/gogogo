package main

import "fmt"

// 值传递实例 —— 交换两个值，就需要用到指针和引用
func swap(a, b int) {
	b, a = a, b
}

// 使用指针和引用
func swap1(a, b *int) {
	*b, *a = *a, *b
}

// 还有一种不用指针和引用的办法
func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	// 这种方式是无法交换的
	m, n := 1, 2
	swap(1, 2)
	fmt.Println("m&n: ", m, n) // m&n:  1 2

	a, b := 3, 4
	swap1(&a, &b)
	fmt.Println("a&b: ", a, b) // a&b:  4 3

	c, d := 5, 6
	c, d = swap2(c, d)
	fmt.Println("c&d: ", c, d) // c&d:  6 5
}

/**
指针
- go语言中的指针不能运算，不像c语言中的指针是可以运算的
- go语言中参数是值传递的，没有引用传递
*/
