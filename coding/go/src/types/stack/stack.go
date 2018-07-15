package stack

// Stack defines a stack data structure.
type Stack struct {
	values []interface{}
}

// Push adds a new object to the top of the stack.
func (s *Stack) Push(value interface{}) {
	s.values = append(s.values, value)
}

// Pop removes an object from the top of the stack.
func (s *Stack) Pop() {
	length := len(s.values)
	if length == 0 {
		panic("DataType[Stack], Operation[Pop], Reason[Empty stack]")
	}

	s.values = s.values[:length-1]
}

// Peek returns an object from the top of the stack.
func (s *Stack) Peek() interface{} {
	length := len(s.values)
	if length == 0 {
		panic("DataType[Stack], Operation[Peek], Reason[Empty stack]")
	}

	return s.values[length-1]
}

// Empty returns true if the stack contains any object. Otherwise false.
func (s *Stack) Empty() bool {
	return s.Size() == 0
}

// Size returns the number of objects in the stack.
func (s *Stack) Size() int {
	return len(s.values)
}
