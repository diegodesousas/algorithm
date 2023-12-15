package queue_test

import (
	"fmt"
	"testing"

	"algorithm/queue"
)

func TestQueue_push(t *testing.T) {
	q := queue.New[int](
		queue.WithMaxSize(4),
	)

	stream := []int{7, 10, 5, 10, 8, 3, 1, 4, 3, 5, 3, 1}

	for _, value := range stream {
		if q.Exists(value) {
			fmt.Println(value)
		}

		q.Push(value)
	}
}
