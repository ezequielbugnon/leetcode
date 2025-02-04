package structures

type Queue struct {
	data []int
}

func NewQueue() Queue {
	return Queue{
		data: make([]int, 0),
	}
}

func (q *Queue) Enqueu(value int) int {
	q.data = append(q.data, value)
	return len(q.data)
}

func (q *Queue) Dequeu() (int, bool) {
	if len(q.data) <= 0 {
		return 0, false
	}

	value := q.data[0]
	q.data = q.data[1:]

	return value, true

}
