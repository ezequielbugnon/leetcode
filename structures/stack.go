package structures

type Stack struct {
	data []int
}

func New() Stack {
	return Stack{
		data: make([]int, 0),
	}
}

func (s *Stack) Push(value int) int {
	s.data = append(s.data, value)

	return len(s.data)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}

	lastIndex := len(s.data) - 1
	value := s.data[lastIndex]

	s.data = s.data[:lastIndex]

	return value, true
}
