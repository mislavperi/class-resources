package fibonacci

func Generate(n int) int {
	if n <= 1 {
		return n
	}
	return Generate(n-1) + Generate(n-2)
}
