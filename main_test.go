package main

import (
	"testing"
)

func TestSimpleSuccess(t *testing.T) {
	t.Log("Success message")
}

//func TestSimpleFail(t *testing.T) {
//	t.Fail()
//}

//func TestComparisonFail(t *testing.T) {
//	assert.Equal(t, "Expected message", "Actual message")
//}

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
	t.Log("Main test message")

	t.Run("sub test success", func(t *testing.T) {
		t.Log("success message")
	})

	//t.Run("sub test failed", func(t *testing.T) {
	//	t.Log("fail message #2")
	//	t.Fail()
	//})
}
