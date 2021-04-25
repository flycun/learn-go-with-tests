package concurrency

import "testing"

func TestCounterNoSafe(t *testing.T) {

	count := CounterNoSafe()
	t.Logf("result=%d",count)
	expected := 1_000_000
	if count == expected {
		t.Errorf("result epected not eq 1_000_000, count=%d", count)
	}

}
func TestCounterSafe(t *testing.T) {

	count := CounterSafe()
	t.Logf("result=%d",count)
	expected := 1_000_000
	if count != expected {
		t.Errorf("result not epected 1_000_000, count=%d", count)
	}

}


