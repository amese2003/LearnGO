package queue

import "errors"

type Queue struct {
	Items []int
}

func (q *Queue) Push(value int) {
	q.Items = append(q.Items, value)
}

func (q *Queue) Pop() (int, error) {

	if len(q.Items) == 0 {
		return -1, errors.New("NoData")
	}

	item, items := q.Items[0], q.Items[1:]
	q.Items = items
	return item, nil
}

func (q *Queue) Size() int {
	return len(q.Items)
}

func (q *Queue) Peek() (int, error) {
	if len(q.Items) == 0 {
		return -1, errors.New("데이터가 없어요!")
	}

	item := q.Items[0]
	return item, nil
}
