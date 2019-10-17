package usetestgo

import "testing"

func TestFuncs(t *testing.T) {
	s := Hello()
	if s == "hello" {
		t.Log("ok")
	} else {
		t.Fail()
		// t.Log("failed")
	}
}
