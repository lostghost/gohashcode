// Package gohashcode provides a go implementation of the java hashCode method. (http://docs.oracle.com/javase/1.4.2/docs/api/java/lang/String.html#hashCode())
package gohashcode

// New creates a hashcode from the provided string.
func New(in string) int32 {

	// Empty string has a hashcode of 0
	if len(in) == 0 {
		return 0
	}

	// Build hash
	var hash int32
	for _, char := range []byte(in) {
		hash = 31*hash + int32(char)
	}
	return hash
}
