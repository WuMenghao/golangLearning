package mian

import (
	"fmt"
	"unicode/utf8"
)

/**
一、寻找最长不含重复字符串的子串
	思路：(1)记录每个字母最后出现的位置 lastOccurred[x]
		  (2)记录目前的长度 maxlength
		  (3)当lastOccurred[x] >= start 时跟新start
		  (4)当index - start +1 > maxlength时更新 maxlength
          (5)最后返回maxlength

二、rune相当于java的char
1.使用range遍历pos,rune对
2.使用utf8.RuneCountInString获得字符数量
3.使用len获得字节长度
4.使用[]byte获得字节

三、string操作工具包：strings 包
*/

//寻找最长不含重复字符串的子串
func lengthOfNoRepeatingSubStrWithByte(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxlength := 0
	for i, ch := range []byte(s) {

		if lastI, exist := lastOccurred[ch]; exist && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxlength {
			maxlength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	fmt.Println(lastOccurred)
	return maxlength
}

//寻找最长不含重复字符串的子串(uff8版)
func lengthOfNoRepeatingSubStrWithRune(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxlength := 0
	for i, ch := range []rune(s) {

		if lastI, exist := lastOccurred[ch]; exist && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxlength {
			maxlength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	fmt.Println(lastOccurred)
	return maxlength
}

func stringToHex(s string) {
	for _, b := range []byte(s) {
		fmt.Printf("%X \n", b)
	}
}

func stringToChar(s string) {
	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X) ", i, ch)
	}
}

func charToString(chars []rune) {
	for i, ch := range chars { // ch is a rune
		fmt.Printf("(%d %c) ", i, ch)
	}
}

func main() {
	fmt.Println(lengthOfNoRepeatingSubStrWithByte("abcabcbb"))
	fmt.Println(lengthOfNoRepeatingSubStrWithByte("adsfghhja"))
	fmt.Println(lengthOfNoRepeatingSubStrWithByte(""))
	fmt.Println(lengthOfNoRepeatingSubStrWithByte("p"))
	fmt.Println(lengthOfNoRepeatingSubStrWithByte("asdfg"))

	fmt.Println()

	stringToHex("你好我是吴梦昊ABC")
	stringToChar("你好我是吴梦昊ABC")

	//utf8 的 rune 的操作
	fmt.Println(utf8.RuneCountInString("你好我是吴梦昊ABC"))

	fmt.Println()

	fmt.Println(lengthOfNoRepeatingSubStrWithRune("abcabcbb"))
	fmt.Println(lengthOfNoRepeatingSubStrWithRune("adsfghhja"))
	fmt.Println(lengthOfNoRepeatingSubStrWithRune(""))
	fmt.Println(lengthOfNoRepeatingSubStrWithRune("p"))
	fmt.Println(lengthOfNoRepeatingSubStrWithRune("asdfg"))
	fmt.Println(lengthOfNoRepeatingSubStrWithRune("你好我是吴梦昊ABC"))
	fmt.Println(lengthOfNoRepeatingSubStrWithRune("哈啊哈哈啊啊"))

	charToString([]rune("你好我是吴梦昊ABC"))
}
