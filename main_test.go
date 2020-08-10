package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	sm := sum(5, 5)
	if sm != 10 {
		t.Errorf("Fail! sum(5,5) = %d, want 10, got", sm)
	}
}
