package test

import (
	"testing"

	"github.com/dzikuri/exercise/leetcode/romanToInteger"
)

type romanToIntegerTest struct {
	arg1     string
	expected int
}

var romanToIntegerTests = []romanToIntegerTest{
	romanToIntegerTest{arg1: "III", expected: 3},
	romanToIntegerTest{arg1: "IV", expected: 4},
	romanToIntegerTest{arg1: "IX", expected: 9},
	romanToIntegerTest{arg1: "LVIII", expected: 58},
	romanToIntegerTest{arg1: "MCMXCIV", expected: 1994},
}

func Test_RomanToInteger_1(t *testing.T) {
	got := romanToInteger.RomanToInteger("III")
	want := 3
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test_RomanToInteger_2(t *testing.T) {
	for _, test := range romanToIntegerTests {
		if output := romanToInteger.RomanToInteger(test.arg1); output != test.expected {
			t.Errorf("Testing %v Output %v not equal to expected %v", test.arg1, output, test.expected)
		}
	}
}
