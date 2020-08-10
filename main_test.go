package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	sm := sum(4, 5)
	want := 9
	if sm != want {
		t.Errorf("Fail! sm = %d, want %d", sm, want)
	}
	sm = sum(0, 0)
	want = 0
	if sm != want {
		t.Errorf("Fail! sm = %d, want %d", sm, want)
	}
	sm = sum(105,209)
	want = 314 
	if sm != want {
		t.Errorf("Fail! sm = %d, want %d", sm, want)
	}
}
