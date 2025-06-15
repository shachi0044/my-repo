package main

import "testing"

func TestEvenorOdd(t *testing.T) {
	result := EvenorOdd(10)
	if result != "Even" {
		t.Errorf("expected: even, actual: %s", result)
	}
}