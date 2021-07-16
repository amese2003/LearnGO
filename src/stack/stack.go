package stack

import "errors"

type Stack struct {
	Items []int
}

func (q *Stack) Push(value int) {
	q.Items = append(q.Items, value)
}

func (q *Stack) Pop() (int, error) {
	if len(q.Items) == 0 {
		return -1, errors.New("데이터가 없다!")
	}

	item := q.Items[len(q.Items)-1]
	q.Items = q.Items[:len(q.Items)-1]

	return item, nil
}

func (q *Stack) Size() int {
	return len(q.Items)
}

func (q *Stack) Top() (int, error) {
	if len(q.Items) == 0 {
		return -1, errors.New("데이터가 없다!")
	}

	return q.Items[q.Size()-1], nil
}
