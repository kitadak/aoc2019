package main

import (
	"testing"
)

func TestCalculateFuel(t *testing.T) {
	got := CalculateFuel(1969)
	if got != 654 {
		t.Errorf("CalculateFuel(1969) = %d; want 654", got)
	}
	got = CalculateFuel(100756)
	if got != 33583 {
		t.Errorf("CalculateFuel(100756) = %d; want 33583", got)
	}

}
