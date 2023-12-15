package queue

import "errors"

type properties struct {
	maxSize int
}

type Option func(p properties) properties

func WithMaxSize(maxSize int) Option {
	return func(p properties) properties {
		p.maxSize = maxSize
		return p
	}
}

type Queue[T comparable] struct {
	data    []T
	index   map[T]int
	maxSize int
}

func New[T comparable](options ...Option) Queue[T] {
	var p properties
	for _, option := range options {
		p = option(p)
	}

	return Queue[T]{
		data:    make([]T, 0),
		index:   make(map[T]int),
		maxSize: p.maxSize,
	}
}

func (q *Queue[T]) Push(value T) {
	if q.reachedMaximum() {
		_, _ = q.Peek()
	}

	q.data = append(q.data, value)
	q.putIndex(value)

}

func (q *Queue[T]) reachedMaximum() bool {
	return q.maxSize > 0 && len(q.data) == q.maxSize
}

func (q *Queue[T]) Peek() (T, error) {
	var value T

	if len(q.data) == 0 {
		return value, errors.New("err: queue is empty")
	}

	value = q.data[0]

	q.peekIndex(value)

	q.data = q.data[1:]

	return value, nil
}

func (q *Queue[T]) Exists(value T) bool {
	_, ok := q.index[value]

	return ok
}

func (q *Queue[T]) putIndex(value T) {
	if q.Exists(value) {
		q.index[value]++
		return
	}

	q.index[value] = 1
}

func (q *Queue[T]) peekIndex(value T) {
	count, _ := q.index[value]

	if count == 1 {
		delete(q.index, value)
		return
	}

	q.index[value]--
}
