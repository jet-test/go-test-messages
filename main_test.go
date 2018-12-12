package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoroutines(t *testing.T) {
	t.Run("sub test", func(t *testing.T) {
		assert.Equal(t, "success sub test", "error sub test")
	})
	assert.Equal(t, "success main test", "error main test")
}
