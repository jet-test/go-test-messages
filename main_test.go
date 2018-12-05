package main

import (
	"strconv"
	"testing"
)

func TestGoroutines(t *testing.T) {
	for i := 0; i < 10; i++ {
		var name = "test #" + strconv.Itoa(i)
		t.Run(name, func(t *testing.T) {
			t.Error("start", name)
			Sleep()
			t.Error("end", name)
		})
	}
}
