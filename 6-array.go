package main

import "fmt"

// *** 需要注意的是，go中的数组是值类型
func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
	//	改变数组的值，验证数组是值类型，而不是引用类型
	arr[0] = 100
}

func main() {
	// 一维数组定义数组的方式
	//	1，不初始化，默认元素为0
	var arr1 [5]int

	// 2，使用:定义一个数组，初始化值
	arr2 := [3]int{1, 2, 3}

	// 3，使用...，编译器会自动给我们计算int的数量
	arr3 := [...]int{2, 4, 6, 8, 10}

	fmt.Println(arr1, arr2, arr3) // [0 0 0 0 0] [1 2 3] [2 4 6 8 10]

	// 定义一个二维数组
	var grid [4][5]int
	fmt.Println(grid) // [[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]

	// 遍历数组
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	// 使用range关键字
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	// 如果只需要值，不需要下标，可通过下划线_来省略变量
	for _, v := range arr3 {
		fmt.Println(v)
	}

	// 如果只要下标
	for i := range arr3 {
		fmt.Println(i)
	}

	// 打印数组
	printArray(arr1)
	printArray(arr3)
	//printArray(arr2) // 报错：cannot use arr2 (type [3]int) as type [5]int in argument to printArray
	//	因为printArray函数的入参是需要一个[5]int类型的数组，arr1和arr3都是，而arr2不是，因此报错
	fmt.Println(arr1, arr3) // [0 0 0 0 0] [2 4 6 8 10] arr1和arr3并没有被改变
}

/**
为什么go要用range
- 意义明确，美观
- C++：没有类似的能力
- Java/Python：只能for each value，不能同时获取i，v

数组是值类型
- [10]int 和 [20]int 是不同的类型
- 调用func f(arr [10]int)会 "拷贝" 数组
- 在go语言中一般不直接使用数组，而是使用切片slice

那怎么传递不确定长度的数组参数呢？
使用指针和引用
*/
