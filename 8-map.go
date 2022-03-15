package main

import "fmt"

func main() {

	// map的结构：map[K]V
	m1 := map[string]string{
		"name":    "kevin",
		"age":     "18",
		"address": "beijing",
	}

	// 创建一个空的map
	m2 := make(map[string]int)

	var m3 map[string]int // m3 == nil

	// map的结构：map[K1]map[K2]V
	//m2 := map[string]string{}

	fmt.Printf("m1=%v, m2=%v, m3=%v \n: ", m1, m2, m3) // m1=map[address:beijing age:18 name:kevin], m2=map[], m3=map[]

	// map的遍历
	fmt.Println("Traversing map: ")
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// map的操作
	myName := m1["name"]
	fmt.Println("myName: ", myName)                  // myName:  kevin
	fmt.Println("undefinedProperty: ", m1["school"]) // 会输出Zero Value，这里就是空字符串 undefinedProperty:

	// 使用ok判断一个属性是否存在，存在的时候就会输出true，否则false
	mapVal, ok := m1["mapVal"]
	fmt.Println("属性是否存在：", mapVal, ok) // 属性是否存在：  false

	if mapVal1, ok := m1["mapVal1"]; ok {
		fmt.Println(mapVal1)
	} else {
		fmt.Println("Key dose not exist")
	}

	// 删除一个属性
	fmt.Println("Delete a key")
	fmt.Println(m1["age"]) // "18"
	delete(m1, "age")
	age, isExist := m1["age"]
	fmt.Println(age, isExist) // "" false

	fmt.Println(longestNonRepeatingString("abcabcbb")) // 3
	fmt.Println(longestNonRepeatingString("bbbbb"))    // 1
	fmt.Println(longestNonRepeatingString("pwwkew"))   // 3
	fmt.Println(longestNonRepeatingString(""))         // 0
	fmt.Println(longestNonRepeatingString("b"))        // 1
	fmt.Println(longestNonRepeatingString("abcdef"))   // 6
}

/**
map的操作
- 创建：make(map[string]int)
- 获取元素：m[key]
- key不存在时，获得value类型的初始值
- 用value, ok := m[key] 来判断是否存在key这个属性
- delete：删除一个key
- 遍历：使用range遍历key和value，但是不保证顺序遍历，如需顺序，需要手动对key排序
- 使用len来获取元素的个数

map的key的类型
- map使用哈希表，必须可以比较相等
- 除了slice，map，function的内建类型，其他都可以作为key
- Struct类型不包含上述字段的情况下，也可以作为key
*/

/**
LeetCode算法题：寻找最长不含有重复字符的子串
思路：
对于每一个字母x：
- lastOccurred[x]不存在，或者 < start ——》无需操作
- lastOccurred[x]存在，且 >= start ——》更新start
- 更新lastOccurred[x]，更新maxLength
*/
func longestNonRepeatingString(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		// lastOccurred有可能不存在
		lastI, ok := lastOccurred[ch]

		if ok && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccurred[ch] = i
	}

	return maxLength
}
