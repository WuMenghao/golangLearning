package mian

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

/**
一、条件语句
1.if条件语句不需要括号
2.if条件里可以赋值,if条件里的变量作用域仅在if语句里

二、switch case 语句
1.switch case 自动会break ,除非使用fallthrough
2.switch 后面可以没有表达式

三、for 语句
1.for的条件不需要括号
2.for的条件里可以省略初始条件,结束条件,递增表达式
*/

func ifStatement() {
	const fileName = "abc.txt"
	if contents, err := ioutil.ReadFile(fileName); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s \n", contents)
	}
}

func switchStatement(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "D"
	case score < 80 && score >= 60:
		g = "C"
	case score < 90 && score >= 80:
		g = "B"
	case score <= 100 && score >= 90:
		g = "A"
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	}
	return g
}

func forLoop(n int) string {
	//convertToBein
	result := ""
	if n == 0 {
		return "0"
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	//while类似
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		println("abc")
	}
}

func main() {
	ifStatement()

	fmt.Println(
		switchStatement(100),
		switchStatement(59),
		switchStatement(89),
		switchStatement(10),
	)

	fmt.Println(
		forLoop(5),
		forLoop(13),
		forLoop(12334515),
		forLoop(0),
	)

	printFile("abc.txt")
	//forever()
}
