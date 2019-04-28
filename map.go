package main

import "fmt"

/**
一、定义Map(golang中的map是hashMap)
	1.
	m := map[string]string {
		"name":"WuMenghao",
		"age":"24",
		"gender":"male",
	}

	2.
	m2 := make(map[string]string)

	3.
	var m3 map[string]string

二、遍历Map

	for k , v := range m {
		fmt.Println(k,v)
	}

三、map的操作
1.创建: make(map[string]string)
2.获取:m[key]
3.key不存在时，获取Value类型的初始值
4.用 value, exist := m[key] 来判断是否存在key

四、map的Key
1.map使用哈希表，必须可以比较相等
2.除了slice,map,function的内建类型都可以作为key
3.Struct类型(类？)不包括上述字段，也可以做为key

*/

func makeMap() {
	//map的定义
	m1 := map[string]string{
		"name":   "WuMenghao",
		"age":    "24",
		"gender": "male",
	}
	fmt.Println(m1)

	m2 := make(map[string]string)
	fmt.Println(m2)

	var m3 map[string]string
	fmt.Println(m3)
}

func traversingMap(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func gettingValue(m map[string]string) {
	fmt.Println("gettingValue")
	name, exist := m["name"]
	fmt.Println(name, exist)
	name, exist = m["namc"]
	fmt.Println(name, exist)
	fmt.Println()
}

func printValueOfKey(key string, m map[string]string) {
	fmt.Println("printValueOfKey")
	if name, exist := m[key]; exist {
		fmt.Println(name)
	} else {
		fmt.Println("key dose not exist !")
	}
	fmt.Println()
}

func deleteElement(key string, m map[string]string) {
	fmt.Println("Deleting values")

	val, exi := m[key]
	fmt.Println(val, exi)

	delete(m, key)

	val, exi = m[key]
	fmt.Println(val, exi)
}

func main() {
	m := map[string]string{
		"name":   "WuMenghao",
		"age":    "24",
		"gender": "male",
	}

	//创建
	makeMap()
	//变量
	traversingMap(m)
	//获取
	gettingValue(m)
	//判断
	printValueOfKey("name", m)
	printValueOfKey("nama", m)
	//删除
	deleteElement("name", m)
}
