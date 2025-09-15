package main

import "testing"

func TestSimple(t *testing.T) {
    // Simple test that should always pass
    result := 1 + 1
    expected := 2
    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
}
