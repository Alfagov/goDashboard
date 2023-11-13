package test

// fixedRand is an io.Reader that always returns the same sequence of bytes
type FixedRand struct{}

func (r FixedRand) Read(p []byte) (n int, err error) {
	// Replace this with the byte sequence of your choice
	fixedBytes := []byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	return copy(p, fixedBytes), nil
}
