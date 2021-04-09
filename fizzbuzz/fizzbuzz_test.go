package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var fizzbuzz Fizzbuzz

func init() {
	fizzbuzz = Fizzbuzz{}
}

func TestNumToStr(t *testing.T) {
	tests := map[string]struct {
		in       int
		expected string
	}{
		"1":   {in: 1, expected: "1"},
		"2":   {in: 2, expected: "2"},
		"101": {in: 101, expected: "101"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out := fizzbuzz.convert(test.in)
			assert.Equal(t, test.expected, out)
		})
	}
}

func TestFizz(t *testing.T) {
	tests := map[string]struct {
		in       int
		expected string
	}{
		"3":  {in: 3, expected: "Fizz"},
		"99": {in: 99, expected: "Fizz"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out := fizzbuzz.convert(test.in)
			assert.Equal(t, test.expected, out)
		})
	}
}

func TestBuzz(t *testing.T) {
	assert.Equal(t, "Buzz", fizzbuzz.convert(5))
}

func TestFizzBuzz(t *testing.T) {
	assert.Equal(t, "FizzBuzz", fizzbuzz.convert(15))
}
