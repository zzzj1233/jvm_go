package rt

type MethodStack struct {
	stack []*StackFrame
	len   uint16
}

func (this *MethodStack) TopFrame() *StackFrame {
	if this.len == 0 {
		panic("jvm Stack is empty !")
	}
	return this.stack[this.len-1]
}

func (this *MethodStack) PushFrame(frame *StackFrame) {
	if this.len > 1024 {
		panic("java.lang.StackOverFlowError")
	}
	this.stack[this.len] = frame
	this.len++
}

func (this MethodStack) PopFrame() *StackFrame {
	if this.len == 0 {
		panic("jvm Stack is empty !")
	}
	this.len--
	frame := this.stack[this.len]
	this.stack[this.len] = nil
	return frame
}

func NewMethodStack() *MethodStack {
	return &MethodStack{
		stack: make([]*StackFrame, 1024),
		len:   0,
	}
}

type Thread struct {
	Name  string
	Stack *MethodStack
}

func NewThread(name string) *Thread {
	return &Thread{
		Name:  name,
		Stack: NewMethodStack(),
	}
}
