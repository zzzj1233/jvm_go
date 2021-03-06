package test

import (
	"fmt"
	"math"
	"testing"
)
import "../rt"

func TestLocalVarTable(t *testing.T) {
	table := rt.NewLocalVarTable(7)

	table.PutInt(0, 1233)
	table.PutLong(1, 999)
	table.PutDouble(3, 777)
	table.PutFloat(5, 123.4)
	table.PutRef(6, &rt.Object{})

	fmt.Println(table.GetInt(0))
	fmt.Println(table.GetLong(1))
	fmt.Println(table.GetDouble(3))
	fmt.Println(table.GetFloat(5))
	fmt.Println(table.GetRef(6))
}

func TestOperateStack(t *testing.T) {
	stack := rt.NewOperateStack(10)

	stack.PushInt(1)
	stack.PushFloat(1.1)
	stack.PushLong(11)
	stack.PushDouble(11.1)
	stack.PushObj(nil)

	fmt.Println(stack.PopObj())
	fmt.Println(stack.PopDouble())
	fmt.Println(stack.PopLong())
	fmt.Println(stack.PopFloat())
	fmt.Println(stack.PopInt())
}

func Test2(t *testing.T) {
	a := 1065353216
	println(math.Float32frombits(uint32(a)))
}

type Holder struct {
	items []int
}

func Test3(t *testing.T) {
	items := make([]int, 0, 5)

	items = append(items, -1)
	items = append(items, 0)

	h := &Holder{items: items}

	items = append(items, 1)
	items = append(items, 2)
	items = append(items, 3)

	fmt.Println(h.items)
}

func Test4(t *testing.T) {
	items := make([]int, 0, 2)

	fmt.Println(items)

	_ = append(items, -1)
	_ = append(items, 0)

	fmt.Println(items)

	items = append(items, 1)

	fmt.Println(items)
}
