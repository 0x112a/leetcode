package _x232

type MyQueue struct {
	StrackIn []int
	StrackOut []int

}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.StrackIn = append(this.StrackIn,x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.StrackOut) == 0 {
		this.MoveStack()
	}
	n := len(this.StrackOut)
	x := this.StrackOut[n-1]
	this.StrackOut = this.StrackOut[:n-1]
	return x
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.StrackOut) == 0 {
		this.MoveStack()
	}
	n := len(this.StrackOut)

	return this.StrackOut[n-1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.StrackOut) == 0 && len(this.StrackIn) == 0{
		return true
	}

	return false

}

func (this *MyQueue) MoveStack() {
	for len(this.StrackIn)>0{
		this.StrackOut = append(this.StrackOut,this.StrackIn[len(this.StrackIn)-1])
		this.StrackIn = this.StrackIn[:len(this.StrackIn)-1]
	}
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */