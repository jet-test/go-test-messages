package main

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestGoroutines(t *testing.T) {
	for i := 0; i < 5; i++ {
		var name = "test #" + strconv.Itoa(i)
		t.Run(name, func(t *testing.T) {
			t.Error("start", name)
			Sleep(name)
			for i := 0; i < 5; i++ {
				var name = "sub test #" + strconv.Itoa(i) + " from " + name
				t.Run(name, func(t *testing.T) {
					t.Error("start", name)
					Sleep(name)
					t.Error("end", name)
				})
			}
			t.Error("end", name)
			assert.Equal(t, "success "+name, "error "+name)
		})
	}
}
