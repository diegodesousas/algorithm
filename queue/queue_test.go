package queue_test

import (
	"testing"

	"algorithm/queue"

	"github.com/stretchr/testify/assert"
)

func TestQueue_Push_Peek(t *testing.T) {
	q := queue.New[int]()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	value, err := q.Peek()
	assert.Nil(t, err)
	assert.Equal(t, 1, value)

	value, err = q.Peek()
	assert.Nil(t, err)
	assert.Equal(t, 2, value)

	value, err = q.Peek()
	assert.Nil(t, err)
	assert.Equal(t, 3, value)

	value, err = q.Peek()
	assert.EqualError(t, err, "err: queue is empty")
	assert.Equal(t, 0, value)
}

func TestQueue_MaxSize(t *testing.T) {
	q := queue.New[int](
		queue.WithMaxSize(4),
	)

	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)

	value, err := q.Peek()
	assert.Nil(t, err)
	assert.Equal(t, 2, value)
}

func TestQueue_Exists(t *testing.T) {
	q := queue.New[int]()

	q.Push(1)
	q.Push(2)
	q.Push(1)
	q.Push(3)
	q.Push(2)

	exists := q.Exists(1)
	assert.True(t, exists)

	exists = q.Exists(4)
	assert.False(t, exists)

	_, _ = q.Peek()

	exists = q.Exists(1)
	assert.True(t, exists)

	_, _ = q.Peek()
	_, _ = q.Peek()

	exists = q.Exists(1)
	assert.False(t, exists)
}
