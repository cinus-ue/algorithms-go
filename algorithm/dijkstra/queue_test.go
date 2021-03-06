package dijkstra

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue()

	if isEmpty := q.IsEmpty(); !isEmpty {
		t.Errorf("expected q.IsEmpty to report an empty query")
	}

	q.Set("a", 1)
	q.Set("c", 3)
	q.Set("b", 2)

	if isEmpty := q.IsEmpty(); isEmpty {
		t.Errorf("expected q.IsEmpty to report an non-empty query")
	}

	nKeys := len(q.keys)
	nNodes := len(q.nodes)

	if nKeys != 3 {
		t.Errorf("expected queue to have 3 keys instead got %v", nKeys)
	}
	if nNodes != 3 {
		t.Errorf("expected queue to have 3 nodes instead got %v", nNodes)
	}

	firstKey := q.keys[0]
	if firstKey != "a" {
		t.Errorf("expected first key to be a instead got %v", firstKey)
	}

	nextKey, nextPriority := q.Next()
	if nextKey != "a" {
		t.Errorf("expected next key to be a instead got %v", nextKey)
	}
	if nextPriority != 1 {
		t.Errorf("expected priority of a to be 1 instead got %v", nextPriority)
	}

	bPriority, ok := q.Get("b")
	if !ok {
		t.Errorf("expected key b to exist in the queue")
	}
	if bPriority != 2 {
		t.Errorf("expected node b to have a priority of 2 instead got %v", bPriority)
	}
	if len(q.keys) != 2 {
		t.Errorf("expected q.Get to not mutate the queue")
	}
}

func BenchmarkQueueReadWrite(b *testing.B) {
	q := NewQueue()

	q.Set("a", 10)
	q.Set("b", 5)
	q.Set("z", 3)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q.Set("k", 6)
		q.Next()
	}
}
