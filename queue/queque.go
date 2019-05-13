package queue

type Queue []int

func (queue *Queue) Push(value int) {
	*queue = append(*queue, value)
}

func (queue *Queue) Pop() int {
	head := (*queue)[0]
	*queue = (*queue)[1:]
	return head
}

func (queue *Queue) IsEmpty() bool {
	return len(*queue) == 0
}
