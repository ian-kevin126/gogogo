package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval1(a, b int, op string) (q int, r error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		//panic("unsupported operation: " + op)
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

// 13 / 4  = 4...1
func div(a, b int) (int, int) {
	return a / b, a % b
}

// 函数式编程，apply函数第一个参数是函数，第二个是入参
func apply(op func(int, int) int, a, b int) int {
	// 获取函数的指针
	p := reflect.ValueOf(op).Pointer()

	// 获取函数名称
	opName := runtime.FuncForPC(p).Name()

	// 打印函数的信息
	fmt.Printf("Calling function %s with args"+"(%d, %d) \n", opName, a, b)

	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}

	return s
}

func main() {

	// 通过error就不会让程序的执行中断，用panic会中断执行，可以换成打印error的方式
	if result, err := eval1(3, 4, "x"); err != nil {
		fmt.Println("Error: ", err) // Error:  unsupported operation: x
	} else {
		fmt.Println(result)
	}

	fmt.Println(apply(pow, 3, 4)) // Calling function main.pow with args(3, 4) 81

	fmt.Println(
		eval1(3, 4, "*"), // 12 nil
	)

	fmt.Println(
		sum(1, 2, 3, 4, 5), // 15
	)

	fmt.Println(
		div(13, 3), // 4 1
	)
}

/**
函数要点总结：
- 函数返回多个值时可以起名字
- 仅用于非常简单的函数
- 对于调用者而言没有区别
- go语言没有默认参数，函数重载，只有一个可变参数的特性
*/
