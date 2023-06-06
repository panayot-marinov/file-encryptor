package tests

import (
	"file-encryptor/sources"
	"testing"
)

func TestSum(t *testing.T) {
	got := sources.Sum(2, 2)
	want := 4

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

type sumTest struct {
	arg1, arg2, expected int
}

var sumTests = []sumTest{
	sumTest{2, 3, 5},
	sumTest{4, 8, 12},
	sumTest{6, 9, 15},
	sumTest{3, 10, 13},
}

func TestTableSum(t *testing.T) {
	for _, test := range sumTests {
		if output := sources.Sum(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
