package stack

/*
思路：这题也没那么简单，一个栈在push的时候存储，一个栈在pop时返回，问题在于什么时候转移两个栈的数据，并且第二个栈的数据用完之前不能转移
*/

type MyQueue struct {
	stack []int
	queue []int
}

func Constructor() MyQueue {
	return MyQueue{
		stack: make([]int, 0),
		queue: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.queue) == 0 {
		this.stack2queue()
	}
	elem := this.queue[len(this.queue)-1]
	this.queue = this.queue[:len(this.queue)-1]
	return elem
}

func (this *MyQueue) Peek() int {
	if len(this.queue) == 0 {
		this.stack2queue()
	}
	return this.queue[len(this.queue)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack)+len(this.queue) == 0
}

func (this *MyQueue) stack2queue() {
	for len(this.stack) > 0 {
		this.queue = append(this.queue, this.stack[len(this.stack)-1])
		this.stack = this.stack[:len(this.stack)-1]
	}
}
