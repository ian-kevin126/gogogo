package main

import "fmt"

// 一般我们不直接使用数组，而是使用slice，这样就可以不需要每次都传递固定长度的数组了
func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	//定义切片，切片也是值类型的，但是，官方文档里说的是slice是array的一个view（视图）
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[2:6]
	fmt.Println("arr[2:6]: ", s)      // arr[2:6]:  [2 3 4 5]
	fmt.Println("arr[:6]: ", arr[:6]) // arr[:6]:  [0 1 2 3 4 5]
	fmt.Println("arr[2:]: ", arr[2:]) // arr[2:]:  [2 3 4 5 6 7]
	fmt.Println("arr[:]: ", arr[:])   // arr[:]:  [0 1 2 3 4 5 6 7]

	s1 := arr[2:]
	s2 := arr[:]

	fmt.Println("After updateSlice(s1): ")
	updateSlice(s1)
	fmt.Println(s1)  // [100 3 4 5 6 7]
	fmt.Println(arr) // [0 1 100 3 4 5 6 7]   arr也被改变了

	fmt.Println("After updateSlice(s2): ")
	updateSlice(s2)
	fmt.Println(s2)  // [100 1 100 3 4 5 6 7]
	fmt.Println(arr) // [100 1 100 3 4 5 6 7]   arr也被改变了

	s3 := arr[2:6]
	s4 := s3[3:5]             // s3只有4个元素，能取到5吗？
	fmt.Println("arr: ", arr) // arr:  [100 1 100 3 4 5 6 7]
	fmt.Println("s3: ", s3)   // s3:  [100 3 4 5]
	fmt.Println("s4: ", s4)   // s4:  [5 6]

	// 但是直接取超出切片长度的下标元素则是不允许的。
	//fmt.Println("s3[4]: ", s3[4]) // 	fmt.Println("s4: ", s4)   // s4:  [5 6]
	testSlice()
}

func testSlice() {
	// 基于数组创建一个从 a[start] 到 a[end -1] 的切片
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]

	// 创建了一个长度为4的string数组，并返回一个切片给names
	names := []string{"beijing", "shanghai", "guangzhou", "shenzhen"}

	// 创建长度为10的字符串切片
	str := make([]string, 10)

	// 创建长度为10，容量为20的字符串切片
	str1 := make([]string, 10, 20)

	fmt.Println(a, b, names, str, str1) // [76 77 78 79 80] [77 78 79] [beijing shanghai guangzhou shenzhen] [         ] [         ]

	// 基于数组产生新的切片，a[start:end]为左闭右开
	name1 := names[0:3]
	name2 := names[:3]
	name3 := names[2:]
	name4 := names[:]
	fmt.Println(name1, name2, name3, name4)
	// [beijing shanghai guangzhou] [beijing shanghai guangzhou] [guangzhou shenzhen] [beijing shanghai guangzhou shenzhen]

	// 引用传递：切片本身不包含任何数据，它仅仅是底层数组的一个上层表示，对切片的任何修改都将反映在底层数组上。
	names2 := []string{"beijing", "shanghai", "guangzhou", "shenzhen"}
	name5 := names2[0:3]

	// 引用传递，会同时改变原数组
	name5[2] = "kevin"

	fmt.Println(names2) // [beijing shanghai kevin shenzhen]
	fmt.Println(name5)  // [beijing shanghai kevin]

	//	获取切片的长度和容量
	fmt.Println("name5's length: ", len(name5))   // name5's length:  3
	fmt.Println("name5's capacity: ", cap(name5)) // name5's capacity:  4

	// 切片追加元素
	// 如果切片是建立在数组之上的，而数组本身不能改变长度，那么切片是如何动态改变长度的呢？实际发生的情况是，当新元素通过调用 append 函数
	// 追加到切片末尾时，如果超出了容量，append 内部会创建一个新的数组。并将原有数组的元素被拷贝给这个新的数组，最后返回建立在这个新数组上的切片。
	// 这个新切片的容量是旧切片的二倍。

	// append会判断切片是否有剩余空间，如果没有剩余空间，则会自动扩充两倍空间
	name5 = append(name5, "ian")
	fmt.Println("append element to name5: ", name5) // append element to name5:  [beijing shanghai kevin ian]

	// 切片追加切片，只能是向后追加，不能向前扩展
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food) // food: [potatoes tomatoes brinjal oranges apples]

	// 遍历切片
	for i, v := range name5 {
		fmt.Println(i, v)
	}

	// 复制切片
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	copy(slice2, slice1)        // 只会复制slice1的前3个元素到slice2中
	copy(slice1, slice2)        // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1, slice2) // [1 2 3 4 5] [1 2 3]

	// 切片作为函数参数
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos) // slice before function call [8 7 6]
	subtractOne(nos)                               //function modifies the slice
	fmt.Println("slice after function call", nos)  // slice after function call [6 5 4]

	// 切片内存优化
	countriesNeeded := countries()
	fmt.Println(countriesNeeded) // [USA Singapore Germany]

	// 切片追加测试
	appendSlice()

	sliceOps()

	copySlice()
}

// 切片作为函数参数
func subtractOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

func appendSlice() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}

	// 向slice添加元素：
	// 1-添加元素时如果超越cap大小，系统会重新分配更大的底层数组，原来那个如果不再使用会被垃圾回收机制回收
	// 2-由于值传递的关系，必须接收append的返回值
	// 3-s=append(s,val)
	s1 := append(arr, 8)
	s2 := append(arr, 9)
	s3 := append(arr, 10)

	// append会判断切片是否有剩余空间，如果没有剩余空间，则会自动扩充两倍空间
	fmt.Printf("s1=%d, s2=%d, s3=%d", s1, s2, s3)
	// s1=[0 1 2 3 4 5 6 7 8], s2=[0 1 2 3 4 5 6 7 9], s3=[0 1 2 3 4 5 6 7 10]
	fmt.Println("arr: ", arr) // arr:  [0 1 2 3 4 5 6 7] 底层数组没有变化
}

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d \n", len(s), cap(s))
}

func sliceOps() {
	var s []int // Zero value for slice is nil
	for i := 0; i < 30; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s) // [1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41 43 45 47 49 51 53 55 57 59]
}

// copy slice
func copySlice() {
	s1 := []int{2, 4, 6, 8}
	s2 := make([]int, 16)     // 新建一个长度为16的slice
	s3 := make([]int, 10, 32) // 新建一个长度为10但是容量为32的slice

	fmt.Println("Coping slice")
	fmt.Printf("s1=%d, s2=%d, s3=%d \n", s1, s2, s3) // s1=[2 4 6 8], s2=[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], s3=[0 0 0 0 0 0 0 0 0 0]
	copy(s2, s1)
	// 将s1复制到s2前面的四位上，长度和容量都不会变
	fmt.Println("s2: ", s2, len(s2), cap(s2)) // s2:  [2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0] 16 16

	// 删除元素，系统没有提供内建的api，可以使用切片来做
	// s2[:3] + s2[4:] // 这样就可以拼接出一个新数组，等效于删除了8
	s2 = append(s2[:3], s2[4:]...)
	fmt.Println("Delete a element: ", s2) // Delete a element:  [2 4 6 0 0 0 0 0 0 0 0 0 0 0 0]

	// 模拟数组的shift方法
	fmt.Println("Popping from front: ")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println("s2: ", front, s2) // s2:  2 [4 6 0 0 0 0 0 0 0 0 0 0 0 0]

	// 模拟数组的pop方法
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println("s2: ", tail, s2) // s2:  0 [4 6 0 0 0 0 0 0 0 0 0 0 0]
}

/**
切片内存优化
切片保留对底层数组的引用。只要切片存在于内存中，数组就不能被垃圾回收。假设我们有一个非常大的数组，而我们只需要处理它的一小部分，
为此我们创建这个数组的一个切片，并处理这个切片。这时候数组仍然存在于内存中，因为切片正在引用它。

解决该问题的一个方法是使用 copy 函数来创建该切片的一个拷贝。这样我们就可以使用这个新的切片，原来的数组可以被垃圾回收。
*/
func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	// 数组被切片引用，不能回收
	neededCountries := countries[:len(countries)-2]
	// 创建新的切片
	countriesCpy := make([]string, len(neededCountries))
	// 使用 copy 函数创建切片的拷贝
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	// 返回后数组不再被引用，自动回收
	return countriesCpy
}

/**
切片（slice）是建立在数组之上的更方便，更灵活，更强大的数据结构。切片并不存储任何元素而只是对现有数组的引用。
切片的三要素：
- 指向数组中的开始位置
- 切片的长度，通过内置函数len获得
- 切片的最大容量，通过内置函数cap获得

可以认为切片在内部表示为如下的结构体：
type slice struct {
	// 长度
	Length int
	//容量
	Capacity
	// 指向首元素的指针
	zerothElement *byte
}

arr := [...]int{0,1,2,3,4,5,6,7}
s1 := arr[2:6]
s2 := s1[3:5]

- s1的值为[2,3,4,5], s2的值为[5,6]
- slice可以向后扩展，不可以向前扩展
- s[i]不可以超越len(s)，向后扩展不可以超越底层数组cap(s)

参考资料：
Go语言基础之切片：http://www.jalen-qian.com/p/go%E8%AF%AD%E8%A8%80%E5%9F%BA%E7%A1%80%E4%B9%8B%E5%88%87%E7%89%87/
Go切片进阶：https://studygolang.com/articles/34462
Go切片只需这一篇：https://developer.51cto.com/article/676550.html
http://c.biancheng.net/view/27.html
https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-array-and-slice/
https://learnku.com/articles/38316
https://segmentfault.com/a/1190000039949503
知乎：关于Go切片，看这篇就够了：https://zhuanlan.zhihu.com/p/282096939
*/
