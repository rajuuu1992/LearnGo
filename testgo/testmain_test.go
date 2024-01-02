package main

import (
	"fmt"
	"os"
	"testing"
)

type InputOutput struct {
	a   int
	b   int
	res int
}

// Test to verify domain match (Configured and one in interface) function works fine
func TestAdd(t *testing.T) {

	var tests = []struct {
		a   int
		b   int
		res int
	}{
		{1, 2, 3},
		{1, 20, 21},
		{11, 2, 13},
		{1, 21, 22},
	}
	// Add more inputs above

	for _, test := range tests {
		res := Add(test.a, test.b)

		if res != test.res {
			t.Errorf("Add res not matching Expected = %v, actual = %v", res, test.res)
		}
	}

	fmt.Printf(" Coverage at add = %v", testing.Coverage())
}

// Test to verify domain match (Configured and one in interface) function works fine
func TestMul(t *testing.T) {

	var tests = []struct {
		a   int
		b   int
		res int
	}{
		{1, 2, 2},
		{5, 20, 100},
		{11, 2, 22},
		{5, 21, 105},
	}
	// Add more inputs above

	for _, test := range tests {
		res := mul(test.a, test.b)

		if res != test.res {
			t.Errorf("Add res not matching Expected = %v, actual = %v", res, test.res)
		}
	}

	fmt.Printf(" Coverage at mul = %v", testing.Coverage())
}

func TestMain(t *testing.M) {
	res := t.Run()
	expected := 0.8

	if res == 0 && testing.CoverMode() != "" {
		cov := testing.Coverage()

		if cov < expected {
			fmt.Printf("Unit testing Coverage failed %v%%, Expected %v%%\n\n", cov*100, expected*100)
			res = -1
		}
	}
	os.Exit(res)
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(10, 100000)
	}
}

func BenchmarkAdd2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add2(10, 100000)
	}
}
