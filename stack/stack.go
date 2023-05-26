package stack

import "fmt"

// FiFo
type Stack struct {
	vec []int32
}

func New() *Stack {
	return &Stack{}
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
