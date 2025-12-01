package mathy

func Mod(n int, m int) int {
	r := n % m
	if r < 0 {
		r += m
	}
	return r
}
