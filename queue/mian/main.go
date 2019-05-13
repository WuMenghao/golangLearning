package main

import (
	".."
	"fmt"
)

func main() {
	queue := queue.Queue{1}

	queue.Push(2)
	queue.Push(3)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.Pop())
	fmt.Println(queue.IsEmpty())
}
