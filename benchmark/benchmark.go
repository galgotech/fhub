package benchmark

import (
	"fmt"
)

type CountItem struct {
	Increment func(v int64) (int64, bool)
}

var _ func(int64) (int64, bool) = CountItem{}.Increment

//heddle:stream
func Count(n int64, arm CountItem) (int64, error) {
	if n < 0 {
		return 0, fmt.Errorf("benchmark: count: n must not be negative, got %d", n)
	}
	var v int64
	for i := int64(0); i < n; i++ {
		if r, ok := arm.Increment(v); ok {
			v = r
		}
	}
	return v, nil
}

func Add(v int64) int64 {
	return v + 1
}
