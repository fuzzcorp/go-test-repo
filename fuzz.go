package tutorial

// FuzzerEntrypoint is the method Fuzzbuzz will repeatedly
// run with new tests to try and find bugs in BrokenMethod
// small edit
// another small edit
func FuzzerEntrypoint(Data []byte) int {
	// Convert Data to the type we need:
	testString := string(Data)
	BrokenMethod(testString)
	return 0
}
