type MinStack struct {
    items    []int
    minStack []int
}

func Constructor() MinStack {
    return MinStack{}
}

func (this *MinStack) Push(value int) {
    this.items = append(this.items, value)

    if len(this.minStack) == 0 || value <= this.minStack[len(this.minStack)-1] {
        this.minStack = append(this.minStack, value)
    }
}

func (this *MinStack) Pop() {
    top := this.items[len(this.items)-1]
    this.items = this.items[:len(this.items)-1]

    if top == this.minStack[len(this.minStack)-1] {
        this.minStack = this.minStack[:len(this.minStack)-1]
    }
}

func (this *MinStack) Top() int {
    return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(value);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */