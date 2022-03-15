package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "YES我爱北京"
	fmt.Println("字符串输出：", []byte(s))       // 字符串输出： [89 69 83 230 136 145 231 136 177 229 140 151 228 186 172]
	fmt.Println(utf8.RuneCountInString(s)) // 7 中文算一个字符

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes) //
		bytes = bytes[size:]
		fmt.Printf("%c ", ch) // Y E S 我 爱 北 京
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch) // (0 Y) (1 E) (2 S) (3 我) (4 爱) (5 北) (6 京)
	}
	fmt.Println()

	// 适配中文后的最长不含重复字符的子串
	fmt.Println(longestNonRepeatingString1("abcabcbb"))             // 3
	fmt.Println(longestNonRepeatingString1("bbbbb"))                // 1
	fmt.Println(longestNonRepeatingString1("pwwkew"))               // 3
	fmt.Println(longestNonRepeatingString1(""))                     // 0
	fmt.Println(longestNonRepeatingString1("b"))                    // 1
	fmt.Println(longestNonRepeatingString1("abcdef"))               // 6
	fmt.Println(longestNonRepeatingString1("YES我爱北京爱北京"))           // 7
	fmt.Println(longestNonRepeatingString1("一二三三二一"))               // 3
	fmt.Println(longestNonRepeatingString1("我我我我我"))                // 1
	fmt.Println(longestNonRepeatingString1("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花")) // 8
}

/**
rune 相当于 go 的 char
- 使用range遍历pos，rune对
- 使用utf8.RuneCountInString获得字符数量
- 使用len获得字节长度
- 使用[]byte获得字节
*/

/**
LeetCode算法题：寻找最长不含有重复字符的子串（适配中文字符）
思路：
对于每一个字母x：
- lastOccurred[x]不存在，或者 < start ——》无需操作
- lastOccurred[x]存在，且 >= start ——》更新start
- 更新lastOccurred[x]，更新maxLength
*/
func longestNonRepeatingString1(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	// 把这里的byte改成rune
	//for i, ch := range []byte(s) {
	for i, ch := range []rune(s) {
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
