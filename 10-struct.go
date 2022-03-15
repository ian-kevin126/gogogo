package main

import "fmt"

type treeNode struct {
	// 节点值
	value int
	// 左右指针
	left, right *treeNode
}

// 1，显式定义和命名方法接收者（值传递）
// 为结构struct定义方法，函数名前面的(node treeNode)相当于其他语言的this指针，可以理解为接收者，传的是指，记住：go语言中只有值传递
func (node treeNode) print() {
	fmt.Print(node.value, " ")
}

/**
也可以写成常规的函数

func print() (node treeNode) {
	fmt.Print(node.value, " ")
}

但是在调用的时候就是 print(treeNode)了，而上面的结构方法则是这样调用：treeNode.print()
*/

// 2，使用指针作为方法的接收者（引用传递），只有使用指针才可以改变结构内容，需要注意的是：nil指针也可以调用方法
// 这里setValue前面的接收者就要通过指针来实现引用传递
// 意思就是会直接在源结构上改变，这里就区别print方法中的值传递，就会拷贝一份新的数据去做操作
func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node.")
		return
	}
	node.value = value
}

/**
值接收者 & 指针接收者 ?
- 要改变内容必须使用指针接收者
- 结构过大也考虑使用指针接收者，因为值接收者是拷贝，要考虑一个性能问题
- 一致性：如有指针接收者，最好都是指针接收者

和其他语言的比较
- 值接收者是go语言特有的特性
- 值/指针接收者均可接收值/指针
*/

// 工厂函数
func createNode(value int) *treeNode {
	// 注意：虽然返回的是局部变量的地址（这在C++里面是不合法的），但这在go中是可用的。
	return &treeNode{value: value}
}

// 遍历树
func (node *treeNode) traverse() {
	if node == nil {
		return
	}

	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}

	// 使用new也可以新建一个空的节点
	root.right.left = new(treeNode)

	// 使用自定义的工厂函数
	root.left.right = createNode(2)

	// 使用setValue定义一个节点
	root.right.left.setValue(4)

	root.right.left.print() // 4
	root.print()            // 3
	fmt.Println()

	//nodes := []treeNode{
	//	{value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println("nodes: ", nodes) // nodes:  [{3 <nil> <nil>} {0 <nil> <nil>} {6 <nil> 0xc0000a4018}]

	// 定义一个新的变量pRoot，指向root的引用
	pRoot := &root
	fmt.Println("pRoot", pRoot) // pRoot &{3 0xc0000a4030 0xc0000a4048}
	pRoot.print()               // 3
	pRoot.setValue(200)
	fmt.Println("pRoot", pRoot) // pRoot &{200 0xc00000c048 0xc00000c060}
	pRoot.print()               // 200

	fmt.Println("pRoot1: ")
	var pRoot1 *treeNode
	pRoot1.setValue(200)
	pRoot1 = &root
	pRoot1.setValue(300)
	pRoot1.print() // 300

	fmt.Println("traverse: ")
	root.traverse() // 0 2 300 4 5

}

/**
面向对象：
- go语言仅支持封装，不支持继承和多态
- go语言没有class，只有struct

type TreeNode struct {
	Left, Right *TreeNode
	Value       int
}

func (root *TreeNode) traverse() {
	if root == nil {
		return
	}

	root.Left.traverse()
	fmt.Print(root.Value())
	root.Right.traverse()
}
*/
