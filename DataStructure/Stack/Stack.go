package Stack

import "errors"

type Stack []interface{}

var (
	errStatckEmpty = errors.New("stack is empty")
)

func (s Stack) Size() int {
	return len(s)
}

func (s Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Push(val ...interface{}) {
	*s = append(*s, val...)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errStatckEmpty
	}

	res := (*s)[s.Size()-1]
	*s = (*s)[0 : s.Size()-1]
	return res, nil
}

func NewStack(val ...interface{}) *Stack {
	s := &Stack{}
	s.Push(val...)
	return s
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errStatckEmpty
	}

	return (*s)[s.Size()-1], nil
}

func (s *Stack) Clear() {
	*s = nil
}

func (s *Stack) GetValues() []interface{} {
	out := []interface{}{}
	for _, val := range *s {
		out = append(out, val)
	}
	return out
}
