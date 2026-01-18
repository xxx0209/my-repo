package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
   result := EvenOrOdd(10)
   if result != "enen" {
       t.Errorf("expected: even, actual: %s", result)
   )
}
