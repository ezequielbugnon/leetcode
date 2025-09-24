package datastructures

type Queue struct {
	data []int
}

func (q *Queue) Enqueu(value int) {
	q.data = append(q.data, value)
}

func (q *Queue) Dequeu() {
	if len(q.data) == 0 {
		return
	}
	q.data = q.data[1:]
}

func (q *Queue) Peek() (int, bool) {
	if len(q.data) == 0 {
		return 0, false
	}
	return q.data[0], true
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}
