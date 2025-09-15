package main

import "testing"

func TestSimple(t *testing.T) {
    // Always pass
    if true != true {
        t.Error("This should never fail")
    }
}
