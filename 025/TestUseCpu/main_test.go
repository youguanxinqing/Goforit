package main

import "testing"

func TestHello(t *testing.T) {
	t.Parallel()
	if "zhong" != hello() {
		t.Fail()
	}

	t.Log("complete func hello")
}

func TestCalc(t *testing.T) {
	t.Parallel()
	if 512 != calc() {
		t.Fail()
	}

	t.Log("complete func calc")
}

func BenchmarkHello(b *testing.B) {
	hello()
}

func BenchmarkCalc(b *testing.B) {
	calc()
}
