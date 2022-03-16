package main

import (
	"fmt"
	"gogogo/advanced/interface/queue"
)

func main() {
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)

	fmt.Println(q.Pop()) // 1
	fmt.Println(q.Pop()) // 2

	fmt.Println(q.IsEmpty()) // false
	fmt.Println(q.Pop())     // 3
	fmt.Println(q.IsEmpty()) // true

	q.Push("abc")
	fmt.Println(q.Pop()) // abc
}
