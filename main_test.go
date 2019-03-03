package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSimpleSuccess(t *testing.T) {
	t.Log("Success message")
}

func TestSimpleFail(t *testing.T) {
	t.Error("Error message")
	t.Fail()
}

func TestComparisonFail(t *testing.T) {
	assert.Equal(t, "Expected message", "Actual message")
}

func TestSubTestsSuccess(t *testing.T) {
	t.Log("Main test message")

	t.Run("sub test #1", func(t *testing.T) {
		t.Log("sub message #1")
	})

	t.Run("sub test #2", func(t *testing.T) {
		t.Log("sub message #2")
	})

	t.Run("sub test #3", func(t *testing.T) {
		t.Log("sub message #3")
	})
}

func TestSubTestsFail(t *testing.T) {
	t.Log("Main test first message")

	t.Run("sub test #1", func(t *testing.T) {
		t.Log("first message #1")
		time.Sleep(time.Second)
		t.Log("second message #1")
		t.Fail()
	})

	t.Run("sub test #2", func(t *testing.T) {
		t.Log("first sub message #2")
		time.Sleep(time.Second)
		t.Log("second sub message #2")
		t.Fail()
	})

	t.Log("Main test second message")
	t.Fail()
}
