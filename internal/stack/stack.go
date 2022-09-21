package stack

import (
	"errors"
)

type Stack struct {
	arr []int16
}

func (s *Stack) Push(val int16) error {
	s.arr = append(s.arr, val)
	return nil
}

func (s *Stack) Pop() (int16, error) {
	if len(s.arr) == 0 {
		return 0, errors.New("empty stack")
	}
	val := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return val, nil
}

func New() *Stack {
	return &Stack{arr: make([]int16, 0)}
}
