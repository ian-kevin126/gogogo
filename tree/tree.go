package tree

import "fmt"

type Node struct {
	// 节点值
	Value int
	// 左右指针
	Left, Right *Node
}

// Print
// 1，显式定义和命名方法接收者（值传递）
// 为结构struct定义方法，函数名前面的(node Node)相当于其他语言的this指针，可以理解为接收者，传的是指，记住：go语言中只有值传递
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

/**
也可以写成常规的函数

func print() (node Node) {
	fmt.Print(node.Value, " ")
}

但是在调用的时候就是 print(Node)了，而上面的结构方法则是这样调用：Node.print()
*/

// SetValue
// 2，使用指针作为方法的接收者（引用传递），只有使用指针才可以改变结构内容，需要注意的是：nil指针也可以调用方法
// 这里setValue前面的接收者就要通过指针来实现引用传递
// 意思就是会直接在源结构上改变，这里就区别print方法中的值传递，就会拷贝一份新的数据去做操作
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node.")
		return
	}
	node.Value = value
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

// CreateNode 工厂函数
func CreateNode(value int) *Node {
	// 注意：虽然返回的是局部变量的地址（这在C++里面是不合法的），但这在go中是可用的。
	return &Node{Value: value}
}

// Traverse 遍历树
func (node *Node) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
