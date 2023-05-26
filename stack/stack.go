package stack

import "fmt"

func main() {
	stack := Stack{[]int32{}}

	stack.Push(2)
	stack.Push(4)
	stack.Push(7)
	println("popped: %v", stack.Pop())
	stack.Print()
}

// FiFo
type Stack struct {
	vec []int32
}

func (s *Stack) Pop() int32 {
	vec := s.vec[0]
	s.vec = s.vec[1:]
	return vec
}

func (s *Stack) Push(v int32) {
	s.vec = append(s.vec, v)
}

func (s *Stack) Print() {
	for _, v := range s.vec {
		fmt.Printf("%v - ", v)
	}
}
