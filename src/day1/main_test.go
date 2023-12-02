package main

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "hello", expected: "olleh"},
		{input: "world", expected: "dlrow"},
		{input: "12345", expected: "54321"},
		{input: "", expected: ""},
	}

	for _, test := range tests {
		result := reverse(test.input)
		if result != test.expected {
			t.Errorf("reverse(%s) = %s, expected %s", test.input, result, test.expected)
		}
	}
}

func TestFindFirstDigit(t *testing.T) {
	tests := []struct {
		input    string
		expected rune
		reverse  bool
	}{
		{input: "hello", expected: rune(0)},
		{input: "world", expected: rune(0)},
		{input: "12345", expected: rune('1')},
		{input: "", expected: rune(0)},
		{input: "one", expected: rune('1')},
		{input: "eno", expected: rune(0)},
		{input: "two1nine", expected: rune('2')},
		{input: "ggrbl5cthnzlsbjssixpt", expected: rune('5')},
		{input: "qtwonetwodgbhqmtzf82onefive48", expected: rune('2')},
	}

	for _, test := range tests {
		result := findFirstDigit(test.input)
		if result != test.expected {
			t.Errorf("findFirstDigit(%s) = %s, expected %s", test.input, string(result), string(test.expected))
		}
	}
}

func TestFindLastDigit(t *testing.T) {
	tests := []struct {
		input    string
		expected rune
		reverse  bool
	}{
		{input: "hello", expected: rune(0)},
		{input: "world", expected: rune(0)},
		{input: "12345", expected: rune('5')},
		{input: "", expected: rune(0)},
		{input: "one", expected: rune('1')},
		{input: "two1nine", expected: rune('9')},
		{input: "ggrbl5cthnzlsbjssixpt", expected: rune('6')},
	}

	for _, test := range tests {
		result := findLastDigit(test.input)
		if result != test.expected {
			t.Errorf("findLastDigit(%s) = %s, expected %s", test.input, string(result), string(test.expected))
		}
	}
}
