package mian

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "hello , 我爱慕课网"
	fmt.Println(s)
	fmt.Println(len(s))

	//打印字节
	for _, b := range []byte(s) {
		fmt.Printf("%d ", b)
	}
	fmt.Println()

	//打印16进制码
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	//rune
	for i, ch := range s { //ch is a rune(int32 4字节)
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	//utf8.RuneCountInString(s)
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	//utf8.DecodeRune(bytes)
	bytes := []byte(s)
	for len(bytes) > 0 {
		ru, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", ru)
	}
	fmt.Println()

	//直接使用rune数组
	for i, ru := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ru)
	}
	fmt.Println()
}
