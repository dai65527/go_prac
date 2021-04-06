package mypkg

// IsEven returns true if argument is even number, otherwise not.
func IsEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}
