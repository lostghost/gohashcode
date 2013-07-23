// The hashcode package provides a go implementation of the java hashCode method. (http://docs.oracle.com/javase/1.4.2/docs/api/java/lang/String.html#hashCode())
package hashcode

// Creates a new hashcode from the provided string.
func New(in string) int32 {

	// Initialize output
	var hash int32

	// Empty string has a hashcode of 0
	if len(in) == 0 {
		return hash
	}

	// Convert string into slice of bytes
	b := []byte(in)

	// Build hash
	for i := range b {
		char := b[i]
		hash = ((hash << 5) - hash) + int32(char)
	}

	return hash
}
