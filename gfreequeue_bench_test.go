package gfreequeue

import (
	"testing"

	"github.com/scryner/lfreequeue"
)

func BenchmarkTestGfreequeue(b *testing.B) {
	q := New()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q.Enqueue(1)
		_ = q.Dequeue()
	}
}

func BenchmarkTestLfreequeue(b *testing.B) {
	q := lfreequeue.NewQueue()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q.Enqueue(1)
		_, _ = q.Dequeue()
	}
}
