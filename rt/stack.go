package rt

type Stack struct {
	maxSize uint
	size    uint
	top     *StackFrame
}

func newStack() *Stack {
	return &Stack{
		maxSize: 1024,
	}
}

func (this *Stack) pushFrame(frame *StackFrame) {
	if this.size >= this.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if this.top != nil {
		this.top.Low = frame
	}
	this.top = frame
}

func (this *Stack) popFrame() *StackFrame {
	if this.size == 0 {
		panic("java Stack is empty !")
	}
	this.size--
	top := this.top
	this.top = this.top.Low
	return top
}
