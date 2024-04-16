package stringzilla_test

import (
	"testing"

	stringzilla "github.com/rudolfolah/StringZilla"
)

func TestStringzillaToUpper(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"hello", "HELLO"},
		{"Hello", "HELLO"},
		{"HELLO", "HELLO"},
		{"world", "WORLD"},
	}

	for _, test := range tests {
		output := stringzilla.ToUpper(test.input)
		if output != test.expected {
			t.Errorf("Test failed: %s inputted, %s expected, got %s", test.input, test.expected, output)
		} else {
			t.Logf("Test passed: %s inputted, %s expected, got %s", test.input, test.expected, output)
		}
	}
}

func TestStringzillaFind(t *testing.T) {
	var tests = []struct {
		haystack      string
		needle        string
		expectedFound bool
		expected      uint64
	}{
		{"Hello", "l", true, 2},
		{"hello", "e", true, 1},
		{"HELLO", "o", true, 4},
		{"word", "d", true, 3},
	}

	for _, test := range tests {
		found, output := stringzilla.Find(test.haystack, test.needle)
		if found != test.expectedFound || output != test.expected {
			t.Errorf("Test failed: %s inputted, %s searched, got %t, %d", test.haystack, test.needle, test.expectedFound, test.expected)
		} else {
			t.Logf("Test passed: %s inputted, %s searched, got %t, %d", test.haystack, test.needle, test.expectedFound, test.expected)
		}
	}
}
