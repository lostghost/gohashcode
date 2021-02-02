package gohashcode

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		input    string
		expected int32
	}{
		{input: "", expected: 0},
		{input: "banana", expected: -1396355227},
		{input: "foobar", expected: -1268878963},
		{input: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer quis lectus ut eros aliquam elementum sed et nunc. Donec eu auctor nisl. Donec sapien tellus, dignissim id diam sed, porttitor aliquam tellus. Aenean ornare cursus tincidunt. Vivamus non varius odio. Pellentesque feugiat leo lectus, sit amet iaculis sem vulputate et. Aenean non molestie quam, vel aliquet ante. Nam suscipit imperdiet rhoncus. Donec suscipit sodales lectus ut vestibulum.", expected: 1305961714},
		{input: "abcdef{};:<>", expected: 896648102},
	}

	for testNumber, test := range tests {
		t.Run(fmt.Sprint("TestNumber", testNumber), func(t *testing.T) {
			if result := New(test.input); result != test.expected {
				t.Errorf("Error hashing \"%s\": Expected %d, got %d", test.input, test.expected, result)
			}
		})
	}
}

func BenchmarkNew(b *testing.B) {
	benches := map[string]struct {
		input string
	}{
		"Short":  {input: "banana"},
		"Medium": {input: "This is my sample benchmark string."},
		"Long":   {input: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer quis lectus ut eros aliquam elementum sed et nunc. Donec eu auctor nisl. Donec sapien tellus, dignissim id diam sed, porttitor aliquam tellus. Aenean ornare cursus tincidunt. Vivamus non varius odio. Pellentesque feugiat leo lectus, sit amet iaculis sem vulputate et. Aenean non molestie quam, vel aliquet ante. Nam suscipit imperdiet rhoncus. Donec suscipit sodales lectus ut vestibulum."},
	}

	for name, bench := range benches {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = New(bench.input)
			}
		})
	}
}
