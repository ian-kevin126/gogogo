package main

import (
	"fmt"
	"gogogo/advanced/interface/testing"
)

func getRetriever() retriever {
	// 只在这里告诉你返回哪个类型的Retriever，就不需要更改多处了
	return testing.Retriever{}
}

/**
假设我们有两个团队都写了retriever类，到底用哪个呢？因为go语言是强类型语言，类型检查是非常严格的。
于是，就有了接口 interface，让代码和逻辑是一致的。
*/
type retriever interface {
	Get(string) string
}

func main() {
	var r retriever = getRetriever()
	fmt.Println(r.Get("https://www.imooc.com"))
}

/**
接口：

type Traversal interface {
	Traverse()
}

func main() {
	Traversal := getTraversal()
	Traversal.Traverse()
}

接口的概念：
- 强类型语言：熟悉接口的概念
- 弱类型语言：没（少）有接口的概念
- 接口的详解：使用Google Guice实现依赖注入

Duck typing (鸭子类型)
*/
