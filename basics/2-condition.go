package main

import (
	"fmt"
	"io/ioutil"
)

// if条件语句
func bounded(v int) int {
	if v > 100 {
		return 100
	} else if v < 100 {
		return 0
	} else {
		return v
	}
}

func main() {
	fmt.Println(bounded(98)) // 0

	const filename = "abc.txt"

	// go语言有自带的工具方法读取文件的内容，返回两个值，第一个是文件内容，第二个是读取错误时的返回值
	contents, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err) // open abc.txt: no such file or directory
	} else {
		fmt.Printf("%s\n", contents)
	}

	/**
	上式可简写为，需要注意的是
	 1，if条件语句里面是可以赋值的
	 2，if条件里赋值的变量作用域就在这个if语句里，在if之外是无法访问的。
	*/
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err) // open abc.txt: no such file or directory
	} else {
		fmt.Printf("%s\n", contents)
	}
	//在这里访问contents是不被允许的

	fmt.Println(eval(100, 10, "+")) // 110
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
	) // F F C B A A
}

/**
switch 语句
- switch语句会自动break，除非使用fallthrough
*/
func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator: " + op)
	}

	return result
}

func grade(score int) string {
	g := ""

	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}

	return g
}
