package qw

// X is quote word with interce{} as input and output.
func X(q ...interface{}) []interface{} {
	return q
}

// S is quote word with string as input and output.
func S(q ...string) []string {
	return q
}

// I is quote word with int as input and output.
func I(q ...int) []int {
	return q
}
