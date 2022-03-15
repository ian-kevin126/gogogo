package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
循环语句
- for循环的条件不需要括号，和if一样的
- for的条件里可以省略初始条件，结束条件，递增表达式
*/

func convertTiBinary(n int) string {
	if n == 0 {
		return "0"
	}

	result := ""

	// 2,
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	return result
}

// 读取文件操作
func printFile(filename string) {

	fmt.Println(filename)
	file, err := os.Open(filename)
	fmt.Println(file)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	// 这就相当于其他语言中的while语句了
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {

	// 1, 基础的for循环
	sum := 0
	for i := 1; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum) // 4950

	// 3, 死循环
	//for {
	//	fmt.Println("abc")
	//}

	fmt.Println(
		convertTiBinary(5),  // 101
		convertTiBinary(13), // 1101
	)

	printFile("./basics/abc.txt") // 注意这里的路径，不能直接abc.txt，需要从根文件夹往下取
}

/**
基本语法要点：
1，for，if后面的条件没有括号
2，if条件里也可以定义变量
3，没有while循环
4，switch不需要break，也可以直接switch多个条件
*/
