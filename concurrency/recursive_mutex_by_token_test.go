package concurrency

import (
	"testing"
)

func TestStartTokenLayer(t *testing.T) {
	var token int64 = 1024
	r := &TokenRecursiveMutex{}
	r.Lock(token)
	t.Logf("current recuriion %d ", r.recursion)
	if r.recursion != 1 {
		t.Errorf("current recuriion expected 1,but %d", r.recursion)
	}
	r.Lock(token)
	t.Logf("current recuriion %d ", r.recursion)
	if r.recursion != 2 {
		t.Errorf("current recuriion expected 1,but %d", r.recursion)
	}
	r.Unlock(token)
	if r.recursion != 1 {
		t.Errorf("current recuriion expected 1,but %d", r.recursion)
	}
	r.Unlock(token)
	if r.recursion != 0 {
		t.Errorf("current recuriion expected 0,but %d", r.recursion)
	}
}

