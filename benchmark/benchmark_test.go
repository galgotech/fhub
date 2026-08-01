package benchmark

import (
	"strings"
	"testing"
)

// TestCount drives the stream the way the transpiled count.he arm does — each
// reply is the item's value incremented through Add — and checks the feedback
// loop: item i must carry value i, and the final value is n.
func TestCount(t *testing.T) {
	var calls int64
	got, err := Count(1000, CountItem{Increment: func(v int64) (int64, bool) {
		if v != calls {
			t.Fatalf("item %d carries value %d: the previous reply did not feed back", calls, v)
		}
		calls++
		return Add(v), true
	}})
	if err != nil {
		t.Fatalf("Count(1000): %v", err)
	}
	if got != 1000 {
		t.Errorf("Count(1000) = %d, expected 1000", got)
	}
	if calls != 1000 {
		t.Errorf("the arm ran %d times, expected once per item (1000)", calls)
	}
}

// TestCountZero checks the empty stream: no items, no arm calls, value zero.
func TestCountZero(t *testing.T) {
	got, err := Count(0, CountItem{Increment: func(int64) (int64, bool) {
		t.Fatal("the arm must not run for an empty stream")
		return 0, false
	}})
	if err != nil {
		t.Fatalf("Count(0): %v", err)
	}
	if got != 0 {
		t.Errorf("Count(0) = %d, expected 0", got)
	}
}

// TestCountNegative pins the call's genuinely fallible half: a negative n is
// an error, reported before the arm ever runs.
func TestCountNegative(t *testing.T) {
	_, err := Count(-1, CountItem{Increment: func(int64) (int64, bool) {
		t.Fatal("the arm must not run when the call fails")
		return 0, false
	}})
	if err == nil {
		t.Fatal("Count(-1) must fail")
	}
	if !strings.Contains(err.Error(), "negative") {
		t.Errorf("the error does not name the cause: %v", err)
	}
}

// TestCountSkippedReply pins the missing-reply rule of docs/bindings.md §3.8:
// an item whose arm reports no reply (ok == false) leaves the running value
// unchanged — the host invents neither an error nor a value.
func TestCountSkippedReply(t *testing.T) {
	var i int64
	got, err := Count(10, CountItem{Increment: func(v int64) (int64, bool) {
		i++
		if i%2 == 0 {
			return 0, false
		}
		return Add(v), true
	}})
	if err != nil {
		t.Fatalf("Count(10): %v", err)
	}
	if got != 5 {
		t.Errorf("Count(10) with every second reply missing = %d, expected 5", got)
	}
}

// TestAdd pins the primitive.
func TestAdd(t *testing.T) {
	for _, tt := range []struct{ v, want int64 }{{0, 1}, {41, 42}, {-1, 0}} {
		if got := Add(tt.v); got != tt.want {
			t.Errorf("Add(%d) = %d, expected %d", tt.v, got, tt.want)
		}
	}
}
