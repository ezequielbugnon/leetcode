package datastructures

type Stack struct {
	Data []int
}

func (s *Stack) Push(value int) {
	s.Data = append(s.Data, value)
}

func (s *Stack) Pop() {
	if len(s.Data) == 0 {
		return
	}

	s.Data = s.Data[:len(s.Data)-1]
}

func (s *Stack) Peek() (int, bool) {
	if len(s.Data) == 0 {
		return 0, false
	}

	return s.Data[len(s.Data)-1], true
}
