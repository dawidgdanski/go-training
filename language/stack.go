package language

import "fmt"

func init() {
	var intStack Stack[int]
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	value, ok := intStack.Pop()
	fmt.Println(value, ok)

	var comparableStack ComparableStack[int]
	comparableStack.Push(10)
	comparableStack.Push(20)
	comparableStack.Push(30)
	fmt.Println(comparableStack.Contains(20))
	fmt.Println(comparableStack.Contains(25))
}

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

type ComparableStack[T comparable] struct {
	vals []T
}

func (s *ComparableStack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *ComparableStack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func (s *ComparableStack[T]) Contains(value T) bool {
	for _, val := range s.vals {
		if val == value {
			return true
		}
	}

	return false
}
