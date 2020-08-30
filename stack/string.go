package stack

type StringStack interface {
	Push(string)
	Pop() string
	Size() int
	IsEmpty() bool
}

func NewStringStack() StringStack {
	return &stringStack{}
}

type stringStack struct {
	StringStack

	stack []string
}

func (s *stringStack) Push(data string) {
	s.stack = append(s.stack, data)
}

func (s *stringStack) Pop() string {
	if s.IsEmpty() {
		return ""
	}

	data := s.stack[0]

	s.stack = s.stack[1:]

	return data
}

func (s *stringStack) Size() int {
	return len(s.stack)
}

func (s *stringStack) IsEmpty() bool {
	return s.Size() == 0
}
