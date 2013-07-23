package hashcode

import (
	"testing"
)

// Tests

func TestNew(t *testing.T) {
	testCases := make(map[string]int32)
	testCases["banana"] = -1396355227
	testCases["foobar"] = -1268878963
	testCases["Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer quis lectus ut eros aliquam elementum sed et nunc. Donec eu auctor nisl. Donec sapien tellus, dignissim id diam sed, porttitor aliquam tellus. Aenean ornare cursus tincidunt. Vivamus non varius odio. Pellentesque feugiat leo lectus, sit amet iaculis sem vulputate et. Aenean non molestie quam, vel aliquet ante. Nam suscipit imperdiet rhoncus. Donec suscipit sodales lectus ut vestibulum."] = 1305961714
	testCases["abcdef{};:<>"] = 896648102

	for testString, testHashcode := range testCases {
		hashcode := New(testString)
		if hashcode != testHashcode {
			t.Errorf("Error hashing %s: Expecting %d, got %d", testString, testHashcode, hashcode)
		}
	}
}

// Benchmarks

func BenchmarkNewShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = New("banana")
	}
}

func BenchmarkNewMed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = New("This is my sample benchmark string.")
	}
}

func BenchmarkNewLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = New("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer quis lectus ut eros aliquam elementum sed et nunc. Donec eu auctor nisl. Donec sapien tellus, dignissim id diam sed, porttitor aliquam tellus. Aenean ornare cursus tincidunt. Vivamus non varius odio. Pellentesque feugiat leo lectus, sit amet iaculis sem vulputate et. Aenean non molestie quam, vel aliquet ante. Nam suscipit imperdiet rhoncus. Donec suscipit sodales lectus ut vestibulum.")
	}
}
