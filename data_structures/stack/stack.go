package stack

import "fmt"

type Stack interface {
	Pop() interface{}
	Push(interface{}) int
	Peak() interface{}
	Visualize() string
}

type stack struct {
	data []int
}

func New() Stack {
	return &stack{}
}

func (s *stack) Pop() interface{} {
	lastIndex := len(s.data) - 1
	item := s.data[lastIndex]
	s.data = s.data[:lastIndex]
	return item
}

func (s *stack) Push(item interface{}) int {
	s.data = append(s.data, item.(int))
	return len(s.data)
}

func (s *stack) Peak() interface{} {
	lastIndex := len(s.data) - 1
	item := s.data[lastIndex]
	return item
}

func (s *stack) Visualize() string {
	var res = "["

	for i := 0; i < len(s.data); i++ {
		res += fmt.Sprintf("%v", s.data[i])

		if i+1 < len(s.data) {
			res += ","
		}
	}

	res += "]"
	return res
}
