package queue

import (
	"testing"
)

func TestEnQueue(t *testing.T) {
	q := New()

	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)

	if size := q.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}

	if q.Head().(int) != 1 {
		t.Errorf("wrong head, expected 1 and got %d", q.Head())
	}

	if q.IsEmpty() {
		t.Errorf("IsEmpty should return false")
	}
}

func TestDeQueue(t *testing.T) {
	q := New()

	q.EnQueue(1)
	q.EnQueue(1)
	q.EnQueue(1)

	q.DeQueue()
	if size := q.Size(); size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}

	q.DeQueue()
	if size := q.Size(); size != 1 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}

	q.DeQueue()
	if size := q.Size(); size != 0 {
		t.Errorf("wrong count, expected 0 and got %d", size)
	}

	if q.Head() != nil {
		t.Errorf("wrong head, expected nil and got %v", q.Head())
	}

	if !q.IsEmpty() {
		t.Errorf("IsEmpty should return true")
	}
}
