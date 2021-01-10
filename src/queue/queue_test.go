package queue

import "testing"

func TestEnqueue(t *testing.T) {
	t.Run("Running queue test `enqueue`...", func(t *testing.T) {
		// Initialize a queue
		q := &Queue{}
		q.Enqueue("Test")
		if q.Front().Value != "Test" {
			t.Errorf("Test `enqueue` failed")
		}
	})
}

func TestDequeue(t *testing.T) {
	t.Run("Running queue test `enqueue`...", func(t *testing.T) {
		// Initialize a queue
		q := &Queue{}
		q.PushBack("Test")
		if q.Dequeue().Value != "Test" {
			t.Errorf("Test `enqueue` failed")
		}
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("Running queue test `is-empty`...", func(t *testing.T) {
		// Initialize a queue
		q := &Queue{}
		if q.IsEmpty() != true {
			t.Errorf("Test `is-empty` failed")
		}
	})
}
