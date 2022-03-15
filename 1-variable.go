package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func variable() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	//var a, b int = 3, 4
	//var s string = "abc"

	// 等价于
	var a, b = 3, 4
	var s = "abc"
	fmt.Println(a, b, s)
}

// 还可以使用var()集中定义变量的
var (
	a = 1
	b = true
	c = "abc"
)

//aa := 123  // 这是不允许的，必须要使用var关键字来声明变量
// 类型推断
func variableTypeDeduction() {
	//var a, b, c, d = 3, "abd", true, 4

	// 等价，可以省略var关键字，但是在顶层（函数之外）是不能用冒号:这种定义方法的
	a, b, c, d := 3, "abc", true, 4
	fmt.Println(a, b, c, d) // 3, "abd", true, 4
}

// 用complex来表示欧拉公式：e(iπ) + 1 = 0
func euler() {
	c := 3 + 4i
	// 求|3 + 4i| = 5
	fmt.Println(cmplx.Abs(c))                      // 5
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1) // (0+1.2246467991473515e-16i)
	// 输出 保留小数点后3位的数 e(iπ) + 1 = 0
	fmt.Printf("%.3f\n", cmplx.Pow(math.E, 1i*math.Pi)+1) // (0.000+0.000i)
}

/**
1，内建变量类型
- bool，string
- (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
- byte(8位), rune(类似char,32位)
- float32, float64, (负数)complex64, complex128

2，强制类型转换（go语言没有隐式类型转换）
*/

// 强制类型转换
func variableTypeForceTransfer() {
	a, b := 3, 4

	//c := math.Sqrt(a*a + b*b) 这样是没法通过的
	c := int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c) // 5
}

// 常量 const 当为数值的时候可作为各种类型使用，不需要再向上面那样做强制类型转换
func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c) // abc.txt 5
}

// 枚举类型
func enums() {
	const (
		//cpp    = 0
		//java   = 1
		//python = 2
		//golang = 3
		cpp = iota // iota自增值，和上面的等价
		_
		java
		python
		golang
		javascript
	)
	//fmt.Println(cpp, java, python, golang, javascript) // 0 1 2 3
	fmt.Println(cpp, java, python, golang, javascript) // 0 2 3 4 5

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb) // 1 1024 1048576 1073741824 1099511627776 1125899906842624
}

func main() {
	fmt.Println("hello world")
	variable()
	variableInitialValue()
	variableTypeDeduction()
	euler()
	variableTypeForceTransfer()
	consts()
	enums()
}

/**
变量定义的要点：
1-变量类型写在变量名之后
2-编译器可推测变量类型
3-没有char，只有rune（32位）
4-原生支持复数类型
*/
