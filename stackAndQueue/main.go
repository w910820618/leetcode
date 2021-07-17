package main

import "fmt"

type CQueue struct {
	in  []int
	out []int
}

func Constructor() CQueue {
	return CQueue{
		in:  []int{},
		out: []int{},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.in = append(this.in, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.out) == 0 {
		for len(this.in) > 0 {
			this.out = append(this.out, this.in[0])
			this.in = this.in[1:]
		}
	}

	if len(this.out) != 0 {
		result := this.out[0]
		this.out = this.out[1:]
		return result
	}

	return -1
}

func main() {
	queue := new(CQueue)
	queue.AppendTail(1)
	queue.AppendTail(8)
	queue.AppendTail(20)
	queue.AppendTail(1)
	queue.AppendTail(11)
	queue.AppendTail(2)
	fmt.Println(queue.DeleteHead())
	fmt.Println(queue.DeleteHead())
	fmt.Println(queue.DeleteHead())
	fmt.Println(queue.DeleteHead())
}
