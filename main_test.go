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
}
