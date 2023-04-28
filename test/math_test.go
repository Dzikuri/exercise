package test

import (
	"testing"

	"github.com/dzikuri/exercise/math"
)

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	addTest{arg1: 2, arg2: 3, expected: 5},
	addTest{4, 8, 12},
	addTest{4, 9, 15},
	addTest{2, 10, 13},
}

func TestAdd_3(t *testing.T) {
	for _, test := range addTests {
		if output := math.Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
