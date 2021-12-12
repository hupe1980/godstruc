package queue

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSimplePriority(t *testing.T) {
	q := NewPriorityQueue(func(i, j interface{}) bool {
		return i.(int) > j.(int)
	})

	q.Push(10)
	q.Push(30)
	q.Push(40)
	q.Push(20)

	e1, _ := q.Pop()
	e2, _ := q.Pop()
	e3, _ := q.Pop()
	e4, _ := q.Pop()

	assert.Equal(t, 40, e1)
	assert.Equal(t, 30, e2)
	assert.Equal(t, 20, e3)
	assert.Equal(t, 10, e4)
}

func TestComplexPriority(t *testing.T) {
	type elem struct {
		data      string
		priority  int
		timestamp time.Time
	}

	q := NewPriorityQueue(func(i, j interface{}) bool {
		e1, _ := i.(*elem)
		e2, _ := j.(*elem)

		if e1.priority > e2.priority {
			return true
		}

		if e1.priority == e2.priority && e1.timestamp.Before(e2.timestamp) {
			return true
		}

		return false
	})

	q.Push(&elem{
		data:      "v1",
		priority:  10,
		timestamp: time.Now(),
	})
	q.Push(&elem{
		data:      "v2",
		priority:  20,
		timestamp: time.Now(),
	})
	q.Push(&elem{
		data:      "v3",
		priority:  30,
		timestamp: time.Now(),
	})
	q.Push(&elem{
		data:      "v4",
		priority:  10,
		timestamp: time.Now(),
	})
	q.Push(&elem{
		data:      "v5",
		priority:  20,
		timestamp: time.Now(),
	})
	q.Push(&elem{
		data:      "v6",
		priority:  30,
		timestamp: time.Now(),
	})

	e1, _ := q.Pop()
	e2, _ := q.Pop()
	e3, _ := q.Pop()
	e4, _ := q.Pop()
	e5, _ := q.Pop()
	e6, _ := q.Pop()

	assert.Equal(t, "v3", e1.(*elem).data)
	assert.Equal(t, "v6", e2.(*elem).data)
	assert.Equal(t, "v2", e3.(*elem).data)
	assert.Equal(t, "v5", e4.(*elem).data)
	assert.Equal(t, "v1", e5.(*elem).data)
	assert.Equal(t, "v4", e6.(*elem).data)
}
