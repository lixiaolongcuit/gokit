package slicex

import "testing"

func TestDistinctInt32(t *testing.T) {
	a := []int32{1, 2, 1, 3, 2, 3}
	b := DistinctInt32(a)
	t.Log(b)
}
