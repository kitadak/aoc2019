package main

import (
	"testing"
)

func TestCalculateFuel(t *testing.T) {
	errorMessage := "CalculateFuel(%d) = %d; want %e"
	testFuel := 1969
	expectedFuel := 654
	got := CalculateFuel(testFuel)
	if got != expectedFuel {
		t.Errorf(errorMessage, testFuel, got, expectedFuel)
	}
	testFuel = 100756
	expectedFuel = 33583
	got = CalculateFuel(testFuel)
	if got != expectedFuel {
		t.Errorf(errorMessage, testFuel, got, expectedFuel)
	}
}

func TestCalculateFuelForFuel(t *testing.T) {
	errorMessage := "CalculateFuelForFuel(%d) = %d; want %d"
	testFuel := 1969
	expectedFuel := 966
	got := CalculateFuelForFuel(testFuel)
	if got != expectedFuel {
		t.Errorf(errorMessage, testFuel, got, expectedFuel)
	}

	testFuel = 100756
	expectedFuel = 50346
	got = CalculateFuelForFuel(testFuel)
	if got != expectedFuel {
		t.Errorf(errorMessage, testFuel, got, expectedFuel)
	}
}
