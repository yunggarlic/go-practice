package main

import "testing"

func TestDivide(t *testing.T) {
	_, err: = divide(100.0, 10.0)
	if err != nil {
		t.Error("Should not have thrown error")
	}
}

func TestBadDivide(t *testing.T) {
	_, err: = divide(100.0, 0)
	if err == nil {
		t.Error("Should have thrown error")
	}
}
