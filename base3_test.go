package base3

import "testing"
import "fmt"

func TestIntToBase3(t *testing.T) {
	tests := []struct {
		input    int
		digits   int
		expected string
	}{
		{0, 5, "00000"},
		{1, 5, "00001"},
		{2, 5, "00002"},
		{3, 5, "00010"},
		{4, 5, "00011"},
		{5, 5, "00012"},
		{8, 5, "00022"},
		{9, 5, "00100"},
		{27, 5, "01000"},
	}

	for _, test := range tests {
		result := IntToBase3(test.input, test.digits)
		fmt.Printf("%d -> %s\n", test.input, result)
		if result != test.expected {
			t.Errorf("IntToBase3(%d) = %s; want %s", test.input, result, test.expected)
		}
	}
}
