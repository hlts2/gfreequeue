package gfreequeue

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	q := New()
	if q == nil {
		t.Error("New is nil")
	}
}

func TestDequeu(t *testing.T) {
	tests := []struct {
		values   []int
		length   int
		expected []interface{}
	}{
		{
			values:   []int{1, 2, 3, 4},
			length:   4,
			expected: []interface{}{1, 2, 3, 4},
		},
	}

	for i, test := range tests {
		q := New()
		if q == nil {
			t.Errorf("New is nil")
		}

		for _, value := range test.values {
			q.Enqueue(value)
		}

		got := make([]interface{}, 0, test.length)
		for i := 0; i < test.length; i++ {
			got = append(got, q.Dequeue())
		}

		if !reflect.DeepEqual(test.expected, got) {
			t.Errorf("tests[%d] - Dequeue is wrong. expected: %v, got: %v", i, test.expected, got)
		}

	}
}

func TestIterator(t *testing.T) {
	tests := []struct {
		values   []int
		expected []interface{}
	}{
		{
			values:   []int{1, 2, 3, 4},
			expected: []interface{}{1, 2, 3, 4},
		},
	}

	for i, test := range tests {
		q := New()
		if q == nil {
			t.Errorf("New is nil")
		}

		for _, value := range test.values {
			q.Enqueue(value)
		}

		got := q.Iterator()

		if !reflect.DeepEqual(test.expected, got) {
			t.Errorf("tests[%d] - Iterator is wrong. expected: %v, got: %v", i, test.expected, got)
		}

	}
}

func TestLength(t *testing.T) {
	q := New()
	if q == nil {
		t.Errorf("New is nil")
	}

	for i := 0; i < 100; i++ {
		q.Enqueue(i)
	}
	got := q.Length()

	if 100 != got {
		t.Errorf("Length is wrong. expected: %v, got: %v", 100, got)
	}

	for i := 0; i < 40; i++ {
		_ = q.Dequeue()
	}

	got = q.Length()
	if 60 != got {
		t.Errorf("Length is wrong. expected: %v, got: %v", 60, got)
	}

	for i := 0; i < 70; i++ {
		_ = q.Dequeue()
	}

	got = q.Length()
	if 0 != got {
		t.Errorf("Length is wrong. expected: %v, got: %v", 0, got)
	}
}
