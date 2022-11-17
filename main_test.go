package main

import "testing"

func TestPrendiNumero(t *testing.T) {
	result := returnaNumero2()
	if result != 2 {
		t.Error("Expected", 2, "Got", result)
	}
}
