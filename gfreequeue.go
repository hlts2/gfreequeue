package gfreequeue

import (
	"github.com/hlts2/lock-free"
)

type node struct {
	value interface{}
	next  *node
}

// Queue -
type Queue interface {
	Enqueue(value interface{})
	Dequeue() interface{}
	Length() int
	Iterator() []interface{}
}

type queue struct {
	node   *node
	length int
	lf     lockfree.LockFree
	tail   *node
}

// New initialize new instance
func New() Queue {
	return &queue{
		node:   nil,
		length: 0,
		lf:     lockfree.New(),
		tail:   nil,
	}
}

// Enqueue -
func (q *queue) Enqueue(value interface{}) {
	q.lf.Wait()

	q.enqueue(value)

	q.lf.Signal()
}

func (q *queue) enqueue(value interface{}) {
	n := &node{
		value: value,
		next:  nil,
	}

	q.length++

	if q.node == nil || q.tail == nil {
		q.node = n
		q.tail = n
		return
	}

	q.tail.next = n
	q.tail = n
}

// Dequeue -
func (q *queue) Dequeue() interface{} {
	q.lf.Wait()

	v := q.dequeue()

	q.lf.Signal()

	return v
}

func (q *queue) dequeue() interface{} {
	if q.node == nil {
		return nil
	}

	n := q.node
	q.node = n.next

	if q.node == nil {
		q.tail = nil
		q.length = 0
	} else {
		q.length--
	}

	return n.value
}

func (q *queue) Length() int {
	q.lf.Wait()
	l := q.length
	q.lf.Signal()
	return l
}

// Iterator -
func (q *queue) Iterator() []interface{} {
	q.lf.Wait()
	values := q.iterator()
	q.lf.Signal()
	return values
}

func (q *queue) iterator() []interface{} {
	values := make([]interface{}, 0, q.length)
	for q.length != 0 {
		values = append(values, q.dequeue())
	}
	return values
}
