package structures

import "fmt"

func DemoStacks() {

	// Declare an array of integers with a fixed size of 5

	stack := make(stack, 0) // Create an empty stack
	fmt.Println(stack)
}

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	// FIXME: What do we do if the stack is empty, though?

	l := len(s)
	return s[:l-1], s[l-1]
}
