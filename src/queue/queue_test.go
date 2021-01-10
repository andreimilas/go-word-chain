package queue

import "testing"

func TestEnqueue(t *testing.T) {
	t.Run("Running queue test `enqueue`...", func(t *testing.T) {
		// Initialize a queue and add an element
		q := &Queue{}
		q.Enqueue("Test")
		if q.Front().Value != "Test" {
			t.Errorf("Test `enqueue` failed")
		}
	})
}

func TestDequeue(t *testing.T) {
	t.Run("Running queue test `enqueue`...", func(t *testing.T) {
		// Initialize a queue, add an element (via list methods)
		q := &Queue{}
		q.PushBack("Test")
		// Dequeue
		if q.Dequeue().Value != "Test" {
			t.Errorf("Test `enqueue` failed")
		}
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("Running queue test `is-empty`...", func(t *testing.T) {
		// Initialize a queue and check if it's empty
		q := &Queue{}
		if q.IsEmpty() != true {
			t.Errorf("Test `is-empty` failed")
		}
	})
}
