package main

import "fmt"

// 函数与闭包，返回值是一个func

// 1，常规的函数式编程（闭包）
func adder() func(value int) int {
	sum := 0

	return func(value int) int {
		sum += value
		return sum
	}
}

// 2，正统的函数式编程（不能由状态，只能有常量和函数）：这种写法相较于第一种就比较难读，但是第一种写法的缺点是有一个自由变量
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	adder := adder()

	for i := 0; i < 10; i++ {
		fmt.Println(adder(i))
	}

	adder2 := adder2(0)

	for i := 0; i < 10; i++ {
		var s int
		s, adder2 = adder2(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, s)
	}
}

/**
函数式编程 VS 函数指针
- 函数是一等公民：参数，变量，返回值都可以是函数
- 函数是一等公民——>高阶函数
- 函数——>闭包

"正统的"函数式编程
- 不可变性：不能有状态，只有常量和函数
- 函数只能有一个参数
- 但go语言不会严格规定上述两个限制
*/
