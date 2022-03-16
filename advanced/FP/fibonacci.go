package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 斐波那契数列
func fibonacci() initGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type initGen func() int

// 上面的initGen是一个函数，是一个类型就能实现接口，这就是go语言灵活的地方
func (g initGen) Read(p []byte) (n int, err error) {
	next := g()
	s := fmt.Sprintf("%d\n", next)

	// 大于10000就收
	if next > 10000 {
		return 0, io.EOF
	}

	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()

	//fmt.Println(f()) // 1
	//fmt.Println(f()) // 1
	//fmt.Println(f()) // 2
	//fmt.Println(f()) // 3
	//fmt.Println(f()) // 5
	//fmt.Println(f()) // 8

	printFileContents(f)

}
